CREATE TABLE IF NOT EXISTS waste_Bins (
    bin_ID UUID,
    city VARCHAR(50) NOT NULL,
    "state" VARCHAR(50) NOT NULL,
    "location" geometry(POINT, 4326) NOT NULL,
    location_last_Update TIMESTAMP NOT NULL,
    PRIMARY KEY (bin_ID)
);

CREATE TABLE IF NOT EXISTS users (
    user_ID TEXT PRIMARY KEY,
    spl_Wallet TEXT UNIQUE NOT NULL,
    curr_Balance DOUBLE PRECISION DEFAULT 0,
    total_Earned DOUBLE PRECISION DEFAULT 0,
    profile_Image BYTEA,
    created_At TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS waste_Logs (
    "time_Stamp" TIMESTAMP NOT NULL DEFAULT NOW(),
    glass_Weighed DOUBLE PRECISION,
    metal_Weighed DOUBLE PRECISION,
    plastics_Weighed DOUBLE PRECISION, 
    paper_Weighed DOUBLE PRECISION,
    "location" geometry(POINT, 4326) NOT NULL,
    bin_ID UUID NOT NULL REFERENCES waste_Bins (bin_ID),
    user_ID TEXT NOT NULL REFERENCES users (user_ID)
);

