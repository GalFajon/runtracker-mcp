package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal("Could not load .env; check its format and permissions")
	}

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	server := mcp.NewServer(&mcp.Implementation{Name: "run-mcp", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "listruns", Description: "List the most recent runs according to the provided parameters."}, func(ctx context.Context, req *mcp.CallToolRequest, input ListRunsInput) (*mcp.CallToolResult, Output, error) {
		return ListRuns(ctx, pool, req, input)
	})

	mcp.AddTool(server, &mcp.Tool{Name: "getrun", Description: "Get a run by ID."}, func(ctx context.Context, req *mcp.CallToolRequest, input GetRunInput) (*mcp.CallToolResult, Run, error) {
		return GetRun(ctx, pool, req, input)
	})

	log.Println("Starting server...")

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
