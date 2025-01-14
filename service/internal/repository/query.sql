-- -- name: AddNewAgent :exec
-- INSERT INTO agents (agent_ID, "is", email, city, "state", "location", bin_location_last_Update) 
-- VALUES ($1, $2, $3, $4, $5, ST_SetSRID(ST_MakePoint($6, $7), 4326), $8);

-- -- name: AddNewWasteLog :exec
-- INSERT INTO waste_Logs (waste_Id, created_At, glass_Weighed, metal_Weighed, plastics_Weighed, paper_Weighed, "location", agent_ID, email) 
-- VALUES ($1, $2, $3, $4, $5, $6, ST_SetSRID(ST_MakePoint($7, $8), 4326), $9, $10);

-- -- name: AddNewExchangeHistory :exec
-- INSERT INTO exchange_History (email, amount, operation, created_At) 
-- VALUES ($1, $2, $3::json, $4);

-- -- name: AddNewAgentAndCollector :one
-- WITH new_agent AS (
--     INSERT INTO agents ("is", email, city, "state", "location", bin_location_last_Update) 
--     VALUES ($1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6), 4326), $7)
--     RETURNING agent_ID
-- )
-- , insert_collector AS (
--     INSERT INTO collectors (email, "role", first_Name, last_Name, display_Name, avatar_URL, agent_ID)
--     VALUES ($8, $9, $10, $11, $12, $13, (SELECT agent_ID FROM new_agent))
--     RETURNING agent_ID
-- )
-- SELECT agent_ID FROM new_agent;

-- -- name: RemoveCollector :exec
-- DELETE FROM collectors
-- WHERE email = $1 AND "role" != 'super_collector';


-- name: AddNewAccount :exec
INSERT INTO accounts (
    account_ID, email, hashed_password, first_Name, last_Name, display_Name, avatar_URL, spl_Wallet, curr_Balance, total_Earned, total_glass_waste, total_metal_waste, total_plastics_waste, total_paper_waste, access_Token, access_Token_Expires, id_Token, last_Update_Email, joined
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
);

-- name: UpsertAccount :one
INSERT INTO accounts (
    account_ID, email, first_Name, last_Name, display_Name, avatar_URL, access_Token, access_Token_Expires, id_Token
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) ON CONFLICT (email) DO UPDATE SET
    account_ID = CASE WHEN EXCLUDED.account_ID = '' THEN accounts.account_ID ELSE EXCLUDED.account_ID END,
    first_Name = CASE WHEN EXCLUDED.first_Name = '' THEN accounts.first_Name ELSE EXCLUDED.first_Name END,
    last_Name = CASE WHEN EXCLUDED.last_Name = '' THEN accounts.last_Name ELSE EXCLUDED.last_Name END,
    display_Name = CASE WHEN EXCLUDED.display_Name = '' THEN accounts.display_Name ELSE EXCLUDED.display_Name END,
    avatar_URL = CASE WHEN EXCLUDED.avatar_URL = '' THEN accounts.avatar_URL ELSE EXCLUDED.avatar_URL END,
    access_Token = CASE WHEN EXCLUDED.access_Token = '' THEN accounts.access_Token ELSE EXCLUDED.access_Token END,
    access_Token_Expires = COALESCE(EXCLUDED.access_Token_Expires, accounts.access_Token_Expires),
    id_Token = CASE WHEN EXCLUDED.id_Token = '' THEN accounts.id_Token ELSE EXCLUDED.id_Token END,
    is_new = FALSE
RETURNING *;

-- name: GetAccount :one
SELECT * 
FROM accounts 
WHERE email = $1 
  AND is_Deleted = FALSE;

-- name: UpdateNickName :one
UPDATE accounts
SET display_Name = $1
WHERE email = $2
RETURNING *;


-- -- name: AddNewWasteFromCollector :one
-- INSERT INTO Waste_Logs (
--     glass_Weighed, metal_Weighed, plastics_Weighed, paper_Weighed, "location", funny_Name, agent_ID, email
-- ) VALUES (
--     $1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6), 4326), $7, $8, $9
-- )
-- RETURNING waste_Id;

-- -- name: GetWasteFromCollectorByAgentIDAndDate :many
-- WITH temp_waste_count AS (
--     SELECT COUNT(*) AS count
--     FROM Waste_Logs
--     WHERE agent_ID = $1 AND created_At::date = $2::date
-- )
-- SELECT
--     wl.waste_Id, wl.created_At AS "date", wl.glass_Weighed, wl.metal_Weighed, wl.plastics_Weighed, wl.paper_Weighed, wl."location", wl.agent_ID, 
--     (SELECT count FROM temp_waste_count) AS temp_waste_count
-- FROM 
--     waste_Logs wl
-- WHERE 
--     wl.agent_ID = $1 AND wl.created_At::date = $2::date;

-- -- name: GetAllCollectorsEmailAndAgentID :many
-- SELECT 
--     email, agent_ID, "role"
-- FROM 
--     collectors;

-- name: AddNewBeneficiary :many
WITH inserted_row AS (
  INSERT INTO beneficiaries (b_name, b_phone_no, service_provider, email)
  VALUES ($1, $2, $3, $4)
  RETURNING *
)
SELECT * 
FROM beneficiaries
WHERE email = $4
UNION
SELECT * 
FROM inserted_row
ORDER BY created_at DESC;

-- name: DeleteBeneficiary :many
WITH deleted_row AS (
    DELETE FROM beneficiaries
    WHERE email = $1
      AND b_phone_no = $2
      AND service_provider = $3
    RETURNING *
)
SELECT b_name, b_phone_no, service_provider
FROM beneficiaries
WHERE NOT (
    beneficiaries.email = $1
    AND beneficiaries.b_phone_no = $2
    AND beneficiaries.service_provider = $3
);


-- name: GetBeneficiariesByEmail :many
SELECT *
FROM beneficiaries
WHERE email = $1
ORDER BY created_at DESC;
