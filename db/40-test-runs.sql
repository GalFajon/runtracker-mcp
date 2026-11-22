-- Optional local test data. Only add points to runs created by this script.
BEGIN;
LOCK TABLE runs IN SHARE ROW EXCLUSIVE MODE;

WITH fixtures(route_label, started_at, duration) AS (
    VALUES ('Test River Path', timestamptz '2026-11-18 08:00Z', interval '300 seconds'),
           ('Test River Path', timestamptz '2026-11-20 08:00Z', interval '285.5 seconds'),
           ('Test Park Path',  timestamptz '2026-11-19 09:00Z', interval '360 seconds')
), inserted AS (
    INSERT INTO runs (route_label, started_at)
    SELECT route_label, started_at FROM fixtures f
    WHERE NOT EXISTS (
        SELECT FROM runs r WHERE r.route_label = f.route_label AND r.started_at = f.started_at
    )
    RETURNING id, route_label, started_at
)
INSERT INTO run_points (run_id, segment_id, point_index, recorded_at, position)
SELECT i.id, 0, p, i.started_at + f.duration * (p / 2.0),
       ST_SetSRID(ST_MakePoint(-6.27 + p * 0.005,
           CASE i.route_label WHEN 'Test Park Path' THEN 53.36 ELSE 53.35 END), 4326)::geography
FROM inserted i JOIN fixtures f USING (route_label, started_at)
CROSS JOIN generate_series(0, 2) AS p;

COMMIT;
