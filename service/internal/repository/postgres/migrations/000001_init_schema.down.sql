DROP TABLE IF EXISTS waste_logs;
DROP TABLE IF EXISTS collectors;
DROP TABLE IF EXISTS agent;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS beneficiaries;
DROP TABLE IF EXISTS accounts;

DROP INDEX IF EXISTS accounts_spl_wallet_unique_idx;

DROP TYPE IF EXISTS collectorRole;
DROP TYPE IF EXISTS agentType;