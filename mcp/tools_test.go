package main

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}
	var err error
	testPool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()
	testPool.Close()
	os.Exit(code)
}

func TestGetRun(t *testing.T) {
	ctx := t.Context()
	var id int
	err := testPool.QueryRow(ctx, "SELECT id FROM runs WHERE route_label = $1 AND started_at = $2::text::timestamptz",
		"Test Park Path", "2026-11-19T09:00:00Z").Scan(&id)
	if err != nil {
		t.Fatalf("find seeded park run: %v", err)
	}
	_, run, err := GetRun(ctx, testPool, nil, GetRunInput{RunId: id})
	if err != nil {
		t.Fatal(err)
	}
	if run.RunId != id || run.Geometry["type"] != "MultiLineString" {
		t.Errorf("expected run %d with MultiLineString geometry, got %+v", id, run)
	}
}

func TestGetRunMissingID(t *testing.T) {
	_, _, err := GetRun(t.Context(), testPool, nil, GetRunInput{RunId: -1})
	if !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("expected ErrNoRows, got %v", err)
	}
}

func TestListRuns(t *testing.T) {
	input := ListRunsInput{RouteLabel: "Test River Path", Limit: 1}
	_, out, err := ListRuns(t.Context(), testPool, nil, input)
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Runs) != 1 {
		t.Fatalf("expected one run, got %d", len(out.Runs))
	}
	run := out.Runs[0]
	if run.RouteLabel != input.RouteLabel || run.DurationSeconds != 285.5 || run.Geometry != nil {
		t.Errorf("expected latest seeded river run without geometry, got %+v", run)
	}
}
