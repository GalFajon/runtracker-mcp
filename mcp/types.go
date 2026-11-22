package main

import (
	"time"
)

type Input struct {
	RouteLabel string `json:"route_label" jsonschema:"Name of the route to list runs for."`
	Limit      int    `json:"limit" jsonschema:"How many runs to list, ordered by most recent first."`
}

type Run struct {
	RunId           int            `json:"run_id" jsonschema:"The unique ID of the run."`
	StartedAt       time.Time      `json:"started_at" jsonschema:"The timestamp of when the run began."`
	RouteLabel      string         `json:"route_label" jsonschema:"Name of the route."`
	DistanceM       float64        `json:"distance_m" jsonschema:"Distance of the run in meters."`
	DurationSeconds float64        `json:"duration_seconds" jsonschema:"Duration of the run in seconds"`
	Geometry        map[string]any `json:"geometry,omitempty"`
}

type Output struct {
	Runs []Run `json:"runs" jsonschema:"Runs ordered from most recent to oldest."`
}
