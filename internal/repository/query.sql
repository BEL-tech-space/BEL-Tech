-- name: AddNewWasteBin :exec
INSERT INTO waste_Bins (bin_ID, city, "state", "location", location_last_Update) 
VALUES ($1, $2, $3, ST_SetSRID(ST_MakePoint($4, $5), 4326), $6);

-- name: AddNewUser :exec
INSERT INTO users (user_ID, spl_Wallet, curr_Balance, total_Earned, total_glass_waste, total_metal_waste, total_plastics_waste, total_paper_waste, "password", profile_Image, joined) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
-- name: AddNewWasteLog :exec
INSERT INTO waste_Logs (created_At, glass_Weighed, metal_Weighed, plastics_Weighed, paper_Weighed, "location", bin_ID, user_ID) 
VALUES ($1, $2, $3, $4, $5, ST_SetSRID(ST_MakePoint($6, $7), 4326), $8, $9);

-- name: AddNewExchangeHistory :exec
INSERT INTO exchange_History (user_ID, amount, operation, created_At) 
VALUES ($1, $2, $3::json, $4);

-- name: UpdateUserPassword :exec
UPDATE users SET password = $1 WHERE user_ID = $2;