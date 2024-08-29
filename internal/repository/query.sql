-- name: AddNewWasteBin :exec
INSERT INTO wasteBins (bin_ID, "address", created_At) VALUES ($1, $2, $3);

-- name: AddNewUser :exec
INSERT INTO users (user_ID, spl_Wallet, token_Earned, created_At) VALUES ($1, $2, $3, $4);

-- name: AddNewWasteLog :exec
INSERT INTO wasteLogs ("time_Stamp", glass_Weighed, metal_Weighed, plastics_Weighed, paper_Weighed, bin_ID, user_ID) VALUES ($1, $2, $3, $4, $5, $6, $7);

