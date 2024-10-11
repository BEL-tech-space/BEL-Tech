-- name: AddNewAgent :exec
INSERT INTO agents (agent_ID, "is", email, city, "state", "location", bin_location_last_Update) 
VALUES ($1, $2, $3, $4, $5, ST_SetSRID(ST_MakePoint($6, $7), 4326), $8);

-- name: AddNewUser :exec
INSERT INTO users (
    user_ID, email, hashed_password, first_Name, last_Name, nick_Name, avatar_URL, spl_Wallet, curr_Balance, total_Earned, total_glass_waste, total_metal_waste, total_plastics_waste, total_paper_waste, access_Token, access_Token_Expires, id_Token, last_Update_Email, joined
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
);

-- name: AddNewWasteLog :exec
INSERT INTO waste_Logs (waste_Id, created_At, glass_Weighed, metal_Weighed, plastics_Weighed, paper_Weighed, "location", agent_ID, user_ID) 
VALUES ($1, $2, $3, $4, $5, $6, ST_SetSRID(ST_MakePoint($7, $8), 4326), $9, $10);

-- name: AddNewExchangeHistory :exec
INSERT INTO exchange_History (user_ID, amount, operation, created_At) 
VALUES ($1, $2, $3::json, $4);

-- name: UpsertUser :one
INSERT INTO users (
    user_ID, email, first_Name, last_Name, nick_Name, avatar_URL, access_Token, access_Token_Expires, id_Token
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) ON CONFLICT (email) DO UPDATE SET
    user_ID = COALESCE(EXCLUDED.user_ID, users.user_ID),
    first_Name = CASE WHEN EXCLUDED.first_Name = '' THEN users.first_Name ELSE EXCLUDED.first_Name END,
    last_Name = CASE WHEN EXCLUDED.last_Name = '' THEN users.last_Name ELSE EXCLUDED.last_Name END,
    nick_Name = CASE WHEN EXCLUDED.nick_Name = '' THEN users.nick_Name ELSE EXCLUDED.nick_Name END,
    avatar_URL = CASE WHEN EXCLUDED.avatar_URL = '' THEN users.avatar_URL ELSE EXCLUDED.avatar_URL END,
    access_Token = CASE WHEN EXCLUDED.access_Token = '' THEN users.access_Token ELSE EXCLUDED.access_Token END,
    access_Token_Expires = COALESCE(EXCLUDED.access_Token_Expires, users.access_Token_Expires),
    id_Token = CASE WHEN EXCLUDED.id_Token = '' THEN users.id_Token ELSE EXCLUDED.id_Token END
RETURNING 
    first_Name, last_Name, nick_Name, spl_Wallet, curr_Balance, total_Earned, total_glass_waste, total_metal_waste, total_plastics_waste, total_paper_waste;

-- name: AddNewAgentAndCollector :one
WITH new_agent AS (
    INSERT INTO agents ("is", email, city, "state", "location", bin_location_last_Update) 
    VALUES ($1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6), 4326), $7)
    RETURNING agent_ID
)
, insert_collector AS (
    INSERT INTO collectors (email, "role", first_Name, last_Name, nick_Name, avatar_URL, agent_ID)
    VALUES ($8, $9, $10, $11, $12, $13, (SELECT agent_ID FROM new_agent))
    RETURNING agent_ID
)
SELECT agent_ID FROM new_agent;

-- name: RemoveCollector :exec
DELETE FROM collectors
WHERE email = $1 AND "role" != 'super_collector';

-- name: AddNewTempWasteFromCollector :one
INSERT INTO waste_Logs (
    glass_Weighed, metal_Weighed, plastics_Weighed, paper_Weighed, "location", agent_ID
) VALUES (
    $1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6), 4326), $7
)
RETURNING waste_Id;

-- name: GetWasteFromCollectorByAgentIDAndDate :many
WITH temp_waste_count AS (
    SELECT COUNT(*) AS count
    FROM temp_Waste_Logs
    WHERE agent_ID = $1 AND created_At::date = $2::date
)
SELECT 
    wl.waste_Id, wl.created_At AS "date", wl.glass_Weighed, wl.metal_Weighed, wl.plastics_Weighed, wl.paper_Weighed, wl."location", wl.agent_ID, 
    (SELECT count FROM temp_waste_count) AS temp_waste_count
FROM 
    waste_Logs wl
WHERE 
    wl.agent_ID = $1 AND wl.created_At::date = $2::date;

-- name: GetAllCollectorsEmailAndAgentID :many
SELECT 
    email, agent_ID, "role"
FROM 
    collectors;
