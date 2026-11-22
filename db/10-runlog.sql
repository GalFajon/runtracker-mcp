BEGIN;

CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE runs (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    started_at timestamptz NOT NULL,
    route_label text
);

CREATE INDEX runs_started_at_idx ON runs (started_at DESC, id DESC);

CREATE TABLE run_points (
    run_id bigint NOT NULL REFERENCES runs (id) ON DELETE CASCADE,
    segment_id integer NOT NULL CHECK (segment_id >= 0),
    point_index integer NOT NULL CHECK (point_index >= 0),
    recorded_at timestamptz NOT NULL,
    -- WGS84 coordinates in longitude, latitude order.
    position geography(Point, 4326) NOT NULL,
    elevation_m double precision,
    PRIMARY KEY (run_id, segment_id, point_index)
);


COMMIT;
