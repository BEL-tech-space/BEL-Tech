CREATE TYPE identifyAs AS ENUM ('collector', 'bin');

CREATE TABLE IF NOT EXISTS agents (
    agent_ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "is" identifyAs NOT NULL,
    email TEXT UNIQUE,
    city TEXT,
    "state" TEXT,
    "location" geometry(POINT, 4326),
    bin_location_last_Update TIMESTAMP,
    created_At TIMESTAMP DEFAULT NOW()
);

CREATE TYPE collectorRole AS ENUM ('collector', 'super_collector');

CREATE TABLE IF NOT EXISTS collectors (
    email TEXT PRIMARY KEY,
    "role" collectorRole NOT NULL,
    first_Name VARCHAR(20),
    last_Name VARCHAR(20),
    nick_Name VARCHAR(20),
    avatar_URL TEXT,
    agent_ID UUID REFERENCES agents (agent_ID)
);

CREATE TABLE IF NOT EXISTS temp_Waste_Logs (
    waste_Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_At TIMESTAMP DEFAULT NOW(),
     glass_Weighed REAL NOT NULL DEFAULT 0.0,
    metal_Weighed REAL NOT NULL DEFAULT 0.0,
    plastics_Weighed REAL NOT NULL DEFAULT 0.0, 
    paper_Weighed REAL NOT NULL DEFAULT 0.0,
    "location" geometry(POINT, 4326),
    agent_ID UUID REFERENCES agents (agent_ID)
);

CREATE TABLE IF NOT EXISTS users (
    user_ID TEXT PRIMARY KEY,
    email VARCHAR(64) UNIQUE NOT NULL,
    first_Name VARCHAR(64),
    last_Name VARCHAR(64),
    nick_Name VARCHAR(64),
    avatar_URL TEXT,
    spl_Wallet TEXT UNIQUE,
    curr_Balance DOUBLE PRECISION NOT NULL DEFAULT 0.0,
    total_Earned DOUBLE PRECISION NOT NULL DEFAULT 0.0,
    total_glass_waste REAL NOT NULL DEFAULT 0.0,
    total_metal_waste REAL NOT NULL DEFAULT 0.0,
    total_plastics_waste REAL NOT NULL DEFAULT 0.0,
    total_paper_waste REAL NOT NULL DEFAULT 0.0,
    hashed_password TEXT,
    access_Token TEXT,
    access_Token_Expires TIMESTAMP,
    id_Token TEXT,
    last_Update_Email TIMESTAMP,
    "role" VARCHAR(10) NOT NULL DEFAULT 'user',
    joined TIMESTAMP NOT NULL DEFAULT NOW(),
    is_Deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS waste_Logs (
    waste_Id UUID PRIMARY KEY,
    created_At TIMESTAMP DEFAULT NOW(),
    glass_Weighed REAL NOT NULL DEFAULT 0.0,
    metal_Weighed REAL NOT NULL DEFAULT 0.0,
    plastics_Weighed REAL NOT NULL DEFAULT 0.0, 
    paper_Weighed REAL NOT NULL DEFAULT 0.0,
    "location" geometry(POINT, 4326),
    agent_ID UUID REFERENCES agents(agent_ID) ON DELETE RESTRICT,
    user_ID TEXT REFERENCES users(user_ID) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS exchange_History (
    tran_ID UUID PRIMARY KEY,
    user_ID TEXT NOT NULL REFERENCES users (user_ID),
    amount DOUBLE PRECISION NOT NULL DEFAULT 0.0,
    operation  JSON NOT NULL,
    created_At TIMESTAMP DEFAULT NOW()
);

