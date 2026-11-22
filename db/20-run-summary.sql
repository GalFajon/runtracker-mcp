BEGIN;

CREATE OR REPLACE VIEW run_summary AS
WITH segments AS (
    SELECT run_id, segment_id, max(recorded_at) AS last_point_at,
           CASE WHEN count(*) >= 2 THEN
               ST_MakeLine(position::geometry ORDER BY point_index)
           END AS line
    FROM run_points
    GROUP BY run_id, segment_id
)
SELECT r.id AS run_id, r.started_at, r.route_label,
       extract(epoch FROM max(s.last_point_at) - r.started_at) AS duration_seconds,
       coalesce(sum(ST_Length(s.line::geography)), 0) AS distance_m,
       ST_Collect(s.line ORDER BY s.segment_id)
           FILTER (WHERE s.line IS NOT NULL)::geometry(MultiLineString, 4326) AS geometry
FROM runs AS r
LEFT JOIN segments AS s ON s.run_id = r.id
GROUP BY r.id;

CREATE OR REPLACE VIEW run_distances AS
SELECT run_id, distance_m FROM run_summary;

COMMIT;
