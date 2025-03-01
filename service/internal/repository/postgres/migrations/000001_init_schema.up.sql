CREATE TABLE IF NOT EXISTS accounts (
    account_email VARCHAR(255) PRIMARY KEY,
    account_ID TEXT NOT NULL DEFAULT '',
    first_name VARCHAR(30) NOT NULL DEFAULT '',
    last_name VARCHAR(30) NOT NULL DEFAULT '',
    display_name VARCHAR(30) NOT NULL DEFAULT '',
    avatar_URL TEXT NOT NULL DEFAULT '',
    spl_wallet TEXT NOT NULL DEFAULT '',
    total_bel_point DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_bel_token DOUBLE PRECISION NOT NULL DEFAULT 0,
    bel_point_balance DOUBLE PRECISION NOT NULL DEFAULT 0,
    bel_token_balance DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_glass_waste REAL NOT NULL DEFAULT 0,
    total_metal_waste REAL NOT NULL DEFAULT 0,
    total_plastics_waste REAL NOT NULL DEFAULT 0,
    total_paper_waste REAL NOT NULL DEFAULT 0,
    hashed_password TEXT NOT NULL DEFAULT '',
    hashed_pin TEXT NOT NULL DEFAULT '',
    access_token TEXT NOT NULL DEFAULT '',
    access_token_expires TIMESTAMP,
    id_token TEXT NOT NULL DEFAULT '',
    last_update_email TIMESTAMP,
    "role" VARCHAR(10) NOT NULL DEFAULT 'user',
    joined TIMESTAMP DEFAULT NOW(),
    is_new BOOLEAN NOT NULL DEFAULT TRUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
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
    notif_title TEXT NOT NULL,
    notif_heading TEXT NOT NULL,
    notif_body TEXT NOT NULL,
    days_to_Last INTEGER NOT NULL,
    notif_time TEXT NOT NULL,
    notif_ID TEXT NOT NULL
);


DROP TYPE IF EXISTS collectorRole;
DROP TYPE IF EXISTS identifyAs;

CREATE TYPE collectorRole AS ENUM ('collector', 'super_collector');
CREATE TYPE identifyAs AS ENUM ('collector', 'bin');

CREATE TABLE IF NOT EXISTS agents (
    agent_ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    agent_type identifyAs NOT NULL,
    city TEXT NOT NULL DEFAULT '',
    location_state TEXT NOT NULL DEFAULT '',
    location_addr geometry(POINT, 4326),
    bin_location_last_Update TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS collectors (
    collector_email TEXT PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL DEFAULT '',
    last_name VARCHAR(30) NOT NULL DEFAULT '',
    display_name VARCHAR(30) NOT NULL DEFAULT '',
    avatar_URL TEXT NOT NULL DEFAULT '',
    collector_Role collectorRole NOT NULL DEFAULT 'collector',
    agent_ID UUID REFERENCES agents (agent_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Waste_Logs (
    waste_Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    claimed BOOLEAN DEFAULT FALSE,
    glass_weighed REAL NOT NULL DEFAULT 0,
    metal_weighed REAL NOT NULL DEFAULT 0,
    plastics_weighed REAL NOT NULL DEFAULT 0, 
    paper_weighed REAL NOT NULL DEFAULT 0,
     points_earned DOUBLE PRECISION NOT NULL DEFAULT 0,
    location_addr geometry(POINT, 4326),
    txn_name TEXT NOT NULL DEFAULT '',
    agent_ID UUID REFERENCES agents(agent_ID),
    account_email VARCHAR(255) REFERENCES accounts(account_email)
);




