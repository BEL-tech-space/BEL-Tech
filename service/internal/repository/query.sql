-- name: UpsertAccount :one
INSERT INTO accounts (
    account_ID, account_email, first_name, last_name, display_name, avatar_URL, access_Token, access_token_expires, id_token
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) ON CONFLICT (account_email) DO UPDATE SET
    account_ID = CASE WHEN EXCLUDED.account_ID = '' THEN accounts.account_ID ELSE EXCLUDED.account_ID END,
    first_name = CASE WHEN EXCLUDED.first_name = '' THEN accounts.first_name ELSE EXCLUDED.first_name END,
    last_name = CASE WHEN EXCLUDED.last_name = '' THEN accounts.last_name ELSE EXCLUDED.last_name END,
    display_name = CASE WHEN EXCLUDED.display_name = '' THEN accounts.display_name ELSE EXCLUDED.display_name END,
    avatar_URL = CASE WHEN EXCLUDED.avatar_URL = '' THEN accounts.avatar_URL ELSE EXCLUDED.avatar_URL END,
    access_Token = CASE WHEN EXCLUDED.access_Token = '' THEN accounts.access_Token ELSE EXCLUDED.access_Token END,
    access_token_expires = COALESCE(EXCLUDED.access_Token_Expires, accounts.access_Token_Expires),
    id_token = CASE WHEN EXCLUDED.id_token = '' THEN accounts.id_token ELSE EXCLUDED.id_token END,
    is_new = FALSE
RETURNING *;


-- name: AddNewBeneficiary :many
WITH inserted_row AS (
  INSERT INTO beneficiaries (b_name, b_phone_no, service_provider, account_email)
  VALUES ($1, $2, $3, $4)
  RETURNING *
)
SELECT * 
FROM beneficiaries
WHERE account_email = $4
UNION
SELECT * 
FROM inserted_row
ORDER BY created_at DESC;

-- name: DeleteBeneficiary :many
WITH deleted_row AS (
    DELETE FROM beneficiaries
    WHERE account_email = $1
      AND b_phone_no = $2
      AND service_provider = $3
    RETURNING *
)
SELECT b_name, b_phone_no, service_provider
FROM beneficiaries
WHERE NOT (
    beneficiaries.account_email = $1
    AND beneficiaries.b_phone_no = $2
    AND beneficiaries.service_provider = $3
);


-- name: AddNewAgentAsCollector :one
WITH new_agent AS (
    INSERT INTO agents (agent_type, city, location_state, location_addr) 
    VALUES ($1, $2, $3, ST_SetSRID(ST_MakePoint($4, $5), 4326))
    RETURNING agent_ID
),
insert_collector AS (
    INSERT INTO collectors (collector_email, collector_role, first_name, last_name, display_name, avatar_URL, agent_ID)
    VALUES (
        $6, $7, $8, $9, $10, $11, (SELECT agent_ID FROM new_agent)
    ) ON CONFLICT (collector_email) DO NOTHING
    RETURNING agent_ID
)
SELECT COALESCE(
    (SELECT agent_ID FROM insert_collector), 
    (SELECT agent_ID FROM collectors WHERE collectors.collector_email = $6)
) AS agent_ID;


-- name: AddNewWasteLog :exec
INSERT INTO waste_Logs (waste_Id, created_At, glass_weighed, metal_weighed, plastics_weighed, paper_weighed, location_addr, agent_ID, account_email, points_earned) 
VALUES ($1, $2, $3, $4, $5, $6, ST_SetSRID(ST_MakePoint($7, $8), 4326), $9, $10, $11);


-- name: RemoveCollector :exec
DELETE FROM collectors
WHERE collector_Email = $1 AND collector_Role != 'super_collector';


-- name: GetAccount :one
SELECT * 
FROM accounts 
WHERE account_email = $1 
  AND is_Deleted = FALSE;


-- name: UpdateNickName :one
UPDATE accounts
SET display_name = $1
WHERE account_email = $2
RETURNING *;


-- name: AddNewWasteFromCollector :one
INSERT INTO Waste_Logs (
    glass_weighed, metal_weighed, plastics_weighed, paper_weighed, location_Addr, txn_Name, agent_ID, account_email, points_earned
) VALUES (
    $1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6), 4326), $7, $8, $9, $10
)
RETURNING waste_Id;


-- name: GetAllCollectors :many
SELECT * FROM collectors;


-- name: GetBeneficiariesByEmail :many
SELECT *
FROM beneficiaries
WHERE account_email = $1
ORDER BY created_at DESC;

-- name: CreateNotification :one
INSERT INTO notifications (
    notif_title, notif_heading, notif_body, days_to_Last, notif_time, notif_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id;

-- name: GetAllNotifications :many
SELECT * FROM notifications ORDER BY notif_time DESC;

-- name: DeleteNotificationsByIDs :exec
DELETE FROM notifications WHERE id = ANY($1::int[]);