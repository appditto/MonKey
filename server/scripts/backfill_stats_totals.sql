-- Folds every UTC day before today from stats into the all-time summary tables, one transaction per day.
-- Safe to run live and to rerun: served is recomputed per day, uniques only count rows a run inserted.
-- Run after the rollout, then again after the next UTC midnight to fold in the rollout day. Not right at midnight.
--   psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f scripts/backfill_stats_totals.sql
DO $$
DECLARE
	first_day date := (SELECT (min(created_at) AT TIME ZONE 'UTC')::date FROM stats);
	today date := (now() AT TIME ZONE 'UTC')::date;
	target_day date;
	day_start timestamptz;
	day_end timestamptz;
BEGIN
	IF first_day IS NULL THEN
		RETURN;
	END IF;
	target_day := first_day;
	WHILE target_day < today LOOP
		day_start := target_day::timestamp AT TIME ZONE 'UTC';
		day_end := (target_day + 1)::timestamp AT TIME ZONE 'UTC';

		INSERT INTO stats_service_days AS t (day, service, served)
		SELECT target_day, coalesce(service, ''), sum(count) FROM stats
		WHERE created_at >= day_start AND created_at < day_end
		GROUP BY coalesce(service, '')
		ON CONFLICT (day, service) DO UPDATE SET served = excluded.served;

		WITH inserted AS (
			INSERT INTO stats_addresses (ban_address)
			SELECT DISTINCT ban_address FROM stats
			WHERE created_at >= day_start AND created_at < day_end
			ON CONFLICT DO NOTHING
			RETURNING 1
		)
		INSERT INTO stats_days AS t (day, new_addresses)
		SELECT target_day, count(*) FROM inserted HAVING count(*) > 0
		ON CONFLICT (day) DO UPDATE SET new_addresses = t.new_addresses + excluded.new_addresses;

		WITH inserted AS (
			INSERT INTO stats_clients (ip_address)
			SELECT DISTINCT ip_address FROM stats
			WHERE created_at >= day_start AND created_at < day_end
			ON CONFLICT DO NOTHING
			RETURNING 1
		)
		INSERT INTO stats_days AS t (day, new_clients)
		SELECT target_day, count(*) FROM inserted HAVING count(*) > 0
		ON CONFLICT (day) DO UPDATE SET new_clients = t.new_clients + excluded.new_clients;

		WITH inserted AS (
			INSERT INTO stats_service_addresses (service, ban_address)
			SELECT DISTINCT coalesce(service, ''), ban_address FROM stats
			WHERE created_at >= day_start AND created_at < day_end
			ON CONFLICT DO NOTHING
			RETURNING service
		)
		INSERT INTO stats_service_days AS t (day, service, new_addresses)
		SELECT target_day, service, count(*) FROM inserted GROUP BY service
		ON CONFLICT (day, service) DO UPDATE SET new_addresses = t.new_addresses + excluded.new_addresses;

		COMMIT;
		RAISE NOTICE 'backfilled %', target_day;
		target_day := target_day + 1;
	END LOOP;
END
$$;
