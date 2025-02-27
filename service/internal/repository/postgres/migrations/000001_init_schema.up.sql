CREATE TABLE IF NOT EXISTS accounts (
    email VARCHAR(255) PRIMARY KEY,
    account_ID TEXT NOT NULL DEFAULT '',
    first_Name VARCHAR(30) NOT NULL DEFAULT '',
    last_Name VARCHAR(30) NOT NULL DEFAULT '',
    display_Name VARCHAR(30) NOT NULL DEFAULT '',
    avatar_URL TEXT NOT NULL DEFAULT '',
    spl_Wallet TEXT NOT NULL DEFAULT '',
    curr_Balance DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_Earned DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_glass_waste REAL NOT NULL DEFAULT 0,
    total_metal_waste REAL NOT NULL DEFAULT 0,
    total_plastics_waste REAL NOT NULL DEFAULT 0,
    total_paper_waste REAL NOT NULL DEFAULT 0,
    hashed_password TEXT NOT NULL DEFAULT '',
    access_Token TEXT NOT NULL DEFAULT '',
    access_Token_Expires TIMESTAMP,
    id_Token TEXT NOT NULL DEFAULT '',
    last_Update_Email TIMESTAMP,
    "role" VARCHAR(10) NOT NULL DEFAULT 'user',
    joined TIMESTAMP DEFAULT NOW(),
    is_new BOOLEAN NOT NULL DEFAULT TRUE,
    is_Deleted BOOLEAN NOT NULL DEFAULT FALSE
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
    email VARCHAR(255) REFERENCES accounts(email) ON DELETE CASCADE,
    UNIQUE (b_phone_no, service_provider)
);

CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    notif_title TEXT NOT NULL,
    notif_heading TEXT NOT NULL,
    notif_body TEXT NOT NULL,
    days_to_Last INTEGER NOT NULL,
    notif_time TEXT NOT NULL,
    notif_id TEXT NOT NULL
);


DROP TYPE IF EXISTS collectorRole;
DROP TYPE IF EXISTS identifyAs;

CREATE TYPE collectorRole AS ENUM ('collector', 'super_collector');
CREATE TYPE identifyAs AS ENUM ('collector', 'bin');

CREATE TABLE IF NOT EXISTS agents (
    agent_ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    agent_Type identifyAs NOT NULL,
    email TEXT NOT NULL DEFAULT '',
    city TEXT NOT NULL DEFAULT '',
    location_State TEXT NOT NULL DEFAULT '',
    location_Addr geometry(POINT, 4326),
    bin_location_last_Update TIMESTAMP,
    created_At TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS collectors (
    email TEXT PRIMARY KEY,
    first_Name VARCHAR(30) NOT NULL DEFAULT '',
    last_Name VARCHAR(30) NOT NULL DEFAULT '',
    display_Name VARCHAR(30) NOT NULL DEFAULT '',
    avatar_URL TEXT NOT NULL DEFAULT '',
    collector_Role collectorRole NOT NULL DEFAULT 'collector',
    agent_ID UUID REFERENCES agents (agent_ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Waste_Logs (
    waste_Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_At TIMESTAMP DEFAULT NOW(),
    claimed BOOLEAN DEFAULT FALSE,
    glass_Weighed REAL NOT NULL DEFAULT 0,
    metal_Weighed REAL NOT NULL DEFAULT 0,
    plastics_Weighed REAL NOT NULL DEFAULT 0, 
    paper_Weighed REAL NOT NULL DEFAULT 0,
    location_Addr geometry(POINT, 4326),
    funny_Name TEXT NOT NULL DEFAULT '',
    agent_ID UUID REFERENCES agents(agent_ID),
    email VARCHAR(255) REFERENCES accounts(email)
);




