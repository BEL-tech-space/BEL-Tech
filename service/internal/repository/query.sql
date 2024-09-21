-- name: AddNewWasteBin :exec
INSERT INTO waste_Bins (bin_ID, city, "state", "location", location_last_Update) 
VALUES ($1, $2, $3, ST_SetSRID(ST_MakePoint($4, $5), 4326), $6);

-- name: AddNewUser :exec
INSERT INTO users (
    user_ID, email, hashed_password, first_Name, last_Name, nick_Name, avatar_URL, spl_Wallet, curr_Balance, total_Earned, total_glass_waste, total_metal_waste, total_plastics_waste, total_paper_waste, access_Token, access_Token_Expires, id_Token, last_Update_Email, joined
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
);

-- name: AddNewWasteLog :exec
INSERT INTO waste_Logs (created_At, glass_Weighed, metal_Weighed, plastics_Weighed, paper_Weighed, "location", bin_ID, user_ID) 
VALUES ($1, $2, $3, $4, $5, ST_SetSRID(ST_MakePoint($6, $7), 4326), $8, $9);

-- name: AddNewExchangeHistory :exec
INSERT INTO exchange_History (user_ID, amount, operation, created_At) 
VALUES ($1, $2, $3::json, $4);

-- name: UpsertUser :exec
INSERT INTO users (
    user_ID, email, first_Name, last_Name, nick_Name, avatar_URL, spl_Wallet, curr_Balance, total_Earned, total_glass_waste, total_metal_waste, total_plastics_waste, total_paper_waste, hashed_password, access_Token, access_Token_Expires, id_Token, last_Update_Email
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
) ON CONFLICT (email) DO UPDATE SET
    user_ID = COALESCE(EXCLUDED.user_ID, users.user_ID),
    first_Name = COALESCE(EXCLUDED.first_Name, users.first_Name),
    last_Name = COALESCE(EXCLUDED.last_Name, users.last_Name),
    nick_Name = COALESCE(EXCLUDED.nick_Name, users.nick_Name),
    avatar_URL = COALESCE(EXCLUDED.avatar_URL, users.avatar_URL),
    spl_Wallet = COALESCE(EXCLUDED.spl_Wallet, users.spl_Wallet),
    curr_Balance = COALESCE(EXCLUDED.curr_Balance, users.curr_Balance),
    total_Earned = COALESCE(EXCLUDED.total_Earned, users.total_Earned),
    total_glass_waste = COALESCE(EXCLUDED.total_glass_waste, users.total_glass_waste),
    total_metal_waste = COALESCE(EXCLUDED.total_metal_waste, users.total_metal_waste),
    total_plastics_waste = COALESCE(EXCLUDED.total_plastics_waste, users.total_plastics_waste),
    total_paper_waste = COALESCE(EXCLUDED.total_paper_waste, users.total_paper_waste),
    hashed_password = COALESCE(EXCLUDED.hashed_password, users.hashed_password),
    access_Token = COALESCE(EXCLUDED.access_Token, users.access_Token),
    access_Token_Expires = COALESCE(EXCLUDED.access_Token_Expires, users.access_Token_Expires),
    id_Token = COALESCE(EXCLUDED.id_Token, users.id_Token),
    last_Update_Email = COALESCE(EXCLUDED.last_Update_Email, users.last_Update_Email);