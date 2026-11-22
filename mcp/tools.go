package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func GetRun(ctx context.Context, pool *pgxpool.Pool, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	rows, err := pool.Query(ctx, `
	
	`)

	if err != nil {
		return nil, Output{}, err
	}
	defer rows.Close()

	out := Output{Runs: []Run{}}

	for rows.Next() {
		var row Run

		err := rows.Scan(&row.RunId, &row.DistanceM, &row.DurationSeconds, &row.Geometry, &row.RouteLabel, &row.StartedAt)

		if err != nil {
			return nil, Output{}, err
		}

		out.Runs = append(out.Runs, row)
	}

	if err := rows.Err(); err != nil {
		return nil, Output{}, err
	}

	return nil, out, nil
}

func ListRuns(ctx context.Context, pool *pgxpool.Pool, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	rows, err := pool.Query(ctx, `
		SELECT run_id, started_at, route_label, distance_m, duration_seconds
		FROM run_summary
		WHERE route_label = $1 AND duration_seconds IS NOT NULL
		ORDER BY started_at DESC, run_id DESC
		LIMIT $2`, input.RouteLabel, input.Limit)

	if err != nil {
		return nil, Output{}, err
	}
	defer rows.Close()

	out := Output{Runs: []Run{}}

	for rows.Next() {
		var row Run

		err := rows.Scan(&row.RunId, &row.StartedAt, &row.RouteLabel, &row.DistanceM, &row.DurationSeconds)

		if err != nil {
			return nil, Output{}, err
		}

		out.Runs = append(out.Runs, row)
	}

	if err := rows.Err(); err != nil {
		return nil, Output{}, err
	}

	return nil, out, nil
}
