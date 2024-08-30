-- name: AddNewWasteBin :exec
INSERT INTO waste_Bins (bin_ID, city, "state", "location", location_last_Update) 
VALUES ($1, $2, $3, ST_SetSRID(ST_MakePoint($4, $5), 4326), $6);

-- name: AddNewUser :exec
INSERT INTO users (user_ID, spl_Wallet, curr_Balance, total_Earned, profile_Image, created_At) 
VALUES ($1, $2, $3, $4, $5, $6);

-- name: AddNewWasteLog :exec
INSERT INTO waste_Logs ("time_Stamp", glass_Weighed, metal_Weighed, plastics_Weighed, paper_Weighed, "location", bin_ID, user_ID) 
VALUES ($1, $2, $3, $4, $5, ST_SetSRID(ST_MakePoint($6, $7), 4326), $8, $9);