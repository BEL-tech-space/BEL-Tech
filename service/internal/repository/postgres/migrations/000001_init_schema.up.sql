CREATE TABLE IF NOT EXISTS waste_Bins (
    bin_ID UUID,
    city VARCHAR(50) NOT NULL,
    "state" VARCHAR(50) NOT NULL,
    "location" geometry(POINT, 4326) NOT NULL,
    location_last_Update TIMESTAMP,
    PRIMARY KEY (bin_ID)
);

CREATE TABLE IF NOT EXISTS users (
    user_ID TEXT PRIMARY KEY,
    email VARCHAR(64) UNIQUE NOT NULL,
    first_Name VARCHAR(64),
    last_Name VARCHAR(64),
    nick_Name VARCHAR(64),
    avatar_URL TEXT,
    spl_Wallet TEXT UNIQUE,
    curr_Balance DOUBLE PRECISION DEFAULT 0.0,
    total_Earned DOUBLE PRECISION DEFAULT 0.0,
    total_glass_waste DOUBLE PRECISION DEFAULT 0.0,
    total_metal_waste DOUBLE PRECISION DEFAULT 0.0,
    total_plastics_waste DOUBLE PRECISION DEFAULT 0.0,
    total_paper_waste DOUBLE PRECISION DEFAULT 0.0,
    hashed_password TEXT,
    access_Token TEXT,
    access_Token_Expires TIMESTAMP,
    id_Token TEXT,
    last_Update_Email TIMESTAMP,
    joined TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS waste_Logs (
    created_At TIMESTAMP DEFAULT NOW(),
    glass_Weighed DOUBLE PRECISION DEFAULT 0.0,
    metal_Weighed DOUBLE PRECISION DEFAULT 0.0,
    plastics_Weighed DOUBLE PRECISION DEFAULT 0.0, 
    paper_Weighed DOUBLE PRECISION DEFAULT 0.0,
    "location" geometry(POINT, 4326) NOT NULL,
    bin_ID UUID NOT NULL REFERENCES waste_Bins (bin_ID),
    user_ID TEXT NOT NULL REFERENCES users (user_ID)
);

CREATE TABLE IF NOT EXISTS exchange_History (
    tran_ID UUID PRIMARY KEY,
    user_ID TEXT NOT NULL REFERENCES users (user_ID),
    amount DOUBLE PRECISION NOT NULL DEFAULT 0.0,
    operation  JSON NOT NULL,
    created_At TIMESTAMP DEFAULT NOW()
);

