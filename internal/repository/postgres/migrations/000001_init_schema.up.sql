CREATE TABLE IF NOT EXISTS wasteBins (
    bin_ID TEXT,
    "address" TEXT NOT NULL,
    created_At TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (bin_ID)
);

CREATE TABLE IF NOT EXISTS users (
    user_ID TEXT PRIMARY KEY,
    spl_Wallet TEXT UNIQUE NOT NULL,
    token_Earned DOUBLE PRECISION,
    created_At TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS wasteLogs (
    "time_Stamp" TIMESTAMP NOT NULL DEFAULT NOW(),
    glass_Weighed DOUBLE PRECISION,
    metal_Weighed DOUBLE PRECISION,
    plastics_Weighed DOUBLE PRECISION, 
    paper_Weighed DOUBLE PRECISION,
    bin_ID TEXT NOT NULL REFERENCES wasteBins (bin_ID),
    user_ID TEXT NOT NULL REFERENCES users (user_ID)
);

