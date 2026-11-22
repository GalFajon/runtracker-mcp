BEGIN;

ALTER TABLE runs
    ALTER COLUMN route_label SET NOT NULL,
    ADD CONSTRAINT runs_route_label_nonblank CHECK (route_label ~ '[^[:space:]]');

CREATE VIEW listed_runs AS
SELECT run_id, started_at, route_label, distance_m, duration_seconds
FROM run_summary
WHERE duration_seconds IS NOT NULL;

COMMIT;
