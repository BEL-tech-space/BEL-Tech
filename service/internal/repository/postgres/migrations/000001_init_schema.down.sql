-- Drop tables in reverse order of dependencies

DROP TABLE IF EXISTS waste_logs_from_collector;
DROP TABLE IF EXISTS collectors;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS beneficiaries;
DROP TABLE IF EXISTS accounts;

-- Drop the ENUM type last
DROP TYPE IF EXISTS collectorRole;