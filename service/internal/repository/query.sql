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
    last_Update_Email = COALESCE(EXCLUDED.last_Update_Email, users.last_Update_Email)
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
