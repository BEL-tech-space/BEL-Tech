DROP TYPE IF EXISTS collectorRole;
DROP TYPE IF EXISTS identifyAs;

CREATE TYPE collectorRole AS ENUM ('collector', 'super_collector');
CREATE TYPE identifyAs AS ENUM ('collector', 'bin');

CREATE TABLE IF NOT EXISTS agents (
    agent_ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "is" identifyAs NOT NULL,
    email TEXT NOT NULL DEFAULT '',
    city TEXT NOT NULL DEFAULT '',
    "state" TEXT NOT NULL DEFAULT '',
    "location" geometry(POINT, 4326),
    bin_location_last_Update TIMESTAMP,
    created_At TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS collectors (
    email TEXT PRIMARY KEY,
    first_Name VARCHAR(30) NOT NULL DEFAULT '',
    last_Name VARCHAR(30) NOT NULL DEFAULT '',
    nick_Name VARCHAR(30) NOT NULL DEFAULT '',
    avatar_URL TEXT NOT NULL DEFAULT '',
    "role" collectorRole NOT NULL DEFAULT 'collector',
    agent_ID UUID REFERENCES agents (agent_ID)
);

CREATE TABLE IF NOT EXISTS temp_Waste_Logs (
    waste_Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_At TIMESTAMP DEFAULT NOW(),
     glass_Weighed REAL NOT NULL DEFAULT 0,
    metal_Weighed REAL NOT NULL DEFAULT 0,
    plastics_Weighed REAL NOT NULL DEFAULT 0, 
    paper_Weighed REAL NOT NULL DEFAULT 0,
    "location" geometry(POINT, 4326),
    agent_ID UUID REFERENCES agents (agent_ID)
);

CREATE TABLE IF NOT EXISTS users (
    email TEXT PRIMARY KEY,
    user_ID TEXT NOT NULL DEFAULT '',
    first_Name VARCHAR(30) NOT NULL DEFAULT '',
    last_Name VARCHAR(30) NOT NULL DEFAULT '',
    nick_Name VARCHAR(30) NOT NULL DEFAULT '',
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
    is_Deleted BOOLEAN NOT NULL DEFAULT FALSE
);

DROP INDEX IF EXISTS users_spl_wallet_unique_idx;

CREATE UNIQUE INDEX users_spl_wallet_unique_idx 
ON users (spl_Wallet) 
WHERE spl_Wallet <> '';
 
CREATE TABLE IF NOT EXISTS waste_Logs (
    waste_Id UUID PRIMARY KEY,
    created_At TIMESTAMP DEFAULT NOW(),
    glass_Weighed REAL NOT NULL DEFAULT 0,
    metal_Weighed REAL NOT NULL DEFAULT 0,
    plastics_Weighed REAL NOT NULL DEFAULT 0, 
    paper_Weighed REAL NOT NULL DEFAULT 0,
    "location" geometry(POINT, 4326),
    agent_ID UUID REFERENCES agents(agent_ID) ON DELETE RESTRICT,
    email TEXT NOT NULL REFERENCES users(email) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS exchange_History (
    tran_ID UUID PRIMARY KEY,
    email TEXT NOT NULL REFERENCES users (email),
    amount DOUBLE PRECISION NOT NULL DEFAULT 0,
    operation  JSONB NOT NULL DEFAULT '{}',
    created_At TIMESTAMP DEFAULT NOW()
);

