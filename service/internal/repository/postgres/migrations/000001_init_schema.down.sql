-- Drop tables in reverse order of dependencies

DROP TABLE IF EXISTS waste_logs;
DROP TABLE IF EXISTS collectors;
DROP TABLE IF EXISTS agent;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS beneficiaries;
DROP TABLE IF EXISTS accounts;

-- Drop custom types
DROP TYPE IF EXISTS collectorRole;
DROP TYPE IF EXISTS agentType;

-- Drop indexes
DROP INDEX IF EXISTS accounts_spl_wallet_unique_idx;