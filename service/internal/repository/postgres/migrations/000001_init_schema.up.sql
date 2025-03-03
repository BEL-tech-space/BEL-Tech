CREATE TABLE IF NOT EXISTS accounts (
    account_email VARCHAR(255) PRIMARY KEY,
    account_ID TEXT NOT NULL DEFAULT '',
    first_name VARCHAR(30) NOT NULL DEFAULT '',
    last_name VARCHAR(30) NOT NULL DEFAULT '',
    spl_wallet TEXT NOT NULL DEFAULT '',
    display_name VARCHAR(30) NOT NULL DEFAULT '',
    avatar_URL TEXT NOT NULL DEFAULT '',
    total_bel_point DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_bel_token DOUBLE PRECISION NOT NULL DEFAULT 0,
    bel_point_balance DOUBLE PRECISION NOT NULL DEFAULT 0,
    bel_token_balance DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_glass_waste DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_metal_waste DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_plastics_waste DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_paper_waste DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_glass_points_earned  DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_metal_points_earned  DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_plastics_points_earned  DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_paper_points_earned  DOUBLE PRECISION NOT NULL DEFAULT 0,
    hashed_password TEXT NOT NULL DEFAULT '',
    hashed_pin TEXT NOT NULL DEFAULT '',
    access_token TEXT NOT NULL DEFAULT '',
    access_token_expires TIMESTAMP,
    id_token TEXT NOT NULL DEFAULT '',
    last_update_email TIMESTAMP,
    "role" VARCHAR(10) NOT NULL DEFAULT 'user',
    joined TIMESTAMP DEFAULT NOW(),
    is_new BOOLEAN NOT NULL DEFAULT TRUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP DEFAULT NULL
);

DROP INDEX IF EXISTS accounts_spl_wallet_unique_idx;

CREATE UNIQUE INDEX accounts_spl_wallet_unique_idx 
ON accounts (spl_Wallet) 
WHERE spl_Wallet <> '';

CREATE TABLE IF NOT EXISTS beneficiaries (
    b_name VARCHAR(25) NOT NULL,
    b_phone_no CHAR(11) NOT NULL,
    service_provider VARCHAR(15) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    account_email VARCHAR(255) REFERENCES accounts(account_email) ON DELETE CASCADE,
    UNIQUE (b_phone_no, service_provider)
);

CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    notif_title TEXT NOT NULL DEFAULT '',
    notif_heading TEXT NOT NULL DEFAULT '',
    notif_body TEXT NOT NULL DEFAULT '',
    days_to_Last INTEGER NOT NULL DEFAULT 0,
    notif_time TEXT NOT NULL DEFAULT '',
    notif_ID TEXT NOT NULL DEFAULT ''
);


DROP TYPE IF EXISTS collectorRole;
CREATE TYPE collectorRole AS ENUM ('collector', 'super_collector');

CREATE TABLE IF NOT EXISTS collectors (
    agent_ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    collector_email TEXT UNIQUE,
    first_name VARCHAR(30) NOT NULL DEFAULT '',
    last_name VARCHAR(30) NOT NULL DEFAULT '',
    display_name VARCHAR(30) NOT NULL DEFAULT '',
    avatar_URL TEXT NOT NULL DEFAULT '',
    location_city TEXT NOT NULL DEFAULT '',
    location_state TEXT NOT NULL DEFAULT '',
    location_country TEXT NOT NULL DEFAULT '',
    location_addr geometry(POINT, 4326),
    collector_role collectorRole NOT NULL DEFAULT 'collector',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS waste_logs_from_collector (
    waste_Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    claimed BOOLEAN NOT NULL DEFAULT FALSE,
    glass_weighed REAL NOT NULL DEFAULT 0,
    metal_weighed REAL NOT NULL DEFAULT 0,
    plastics_weighed REAL NOT NULL DEFAULT 0, 
    paper_weighed REAL NOT NULL DEFAULT 0,
    glass_points_earned REAL NOT NULL DEFAULT 0,
    metal_points_earned REAL NOT NULL DEFAULT 0,
    plastics_points_earned REAL NOT NULL DEFAULT 0,
    paper_points_earned REAL NOT NULL DEFAULT 0,
    total_points_earned DOUBLE PRECISION NOT NULL DEFAULT 0,
    location_addr geometry(POINT, 4326),
    txn_name TEXT NOT NULL DEFAULT '',
    agent_ID UUID NOT NULL REFERENCES collectors(agent_ID) ON DELETE RESTRICT,
    account_email TEXT NOT NULL REFERENCES accounts(account_email) ON DELETE RESTRICT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

