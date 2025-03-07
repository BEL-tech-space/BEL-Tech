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
SELECT * FROM beneficiaries
WHERE account_email = $4
UNION
SELECT * 
FROM inserted_row
ORDER BY created_at DESC;


-- name: AddNewAgentAsCollector :one
WITH new_agent AS (
    INSERT INTO agent (collector_email, agent_type)
    VALUES ($1, 'collector')
    RETURNING agent_ID
)
INSERT INTO collectors (
    collector_email, first_name, last_name, display_name, avatar_URL,
    location_city, location_state, location_country, location_addr, 
    agent_ID, collector_role
)
VALUES (
    $1, $2, $3, $4, $5, 
    $6, $7, $8, $9, 
    (SELECT agent_ID FROM new_agent), 'collector'
)
RETURNING *;

-- name: AddNewAgentAsSuperCollector :one
WITH new_agent AS (
    INSERT INTO agent (collector_email, agent_type)
    VALUES ($1, 'collector')
    RETURNING agent_ID
)
INSERT INTO collectors (
    collector_email, first_name, last_name, display_name, avatar_URL,
    location_city, location_state, location_country, location_addr, 
    agent_ID, collector_role
)
VALUES (
    $1, $2, $3, $4, $5, 
    $6, $7, $8, $9, 
    (SELECT agent_ID FROM new_agent), 'super_collector'
)
RETURNING *;

-- name: AddNewWasteFromCollector :exec
WITH new_waste AS (
    INSERT INTO waste_logs (
        claimed, glass_weighed, metal_weighed, plastics_weighed, paper_weighed, 
        glass_points_earned, metal_points_earned, plastics_points_earned, paper_points_earned,
        total_points_earned, location_addr, txn_name, status_report, collected_route, agent_ID, account_email
    ) VALUES (
        $1, $2, $3, $4, $5, 
        $6, $7, $8, $9, 
        $10, ST_SetSRID(ST_MakePoint($11, $12), 4326), $13, $14, 'collector', $15, $16
    )
    RETURNING 
        waste_Id, account_email, total_points_earned, 
        glass_weighed, metal_weighed, plastics_weighed, paper_weighed,
        glass_points_earned, metal_points_earned, plastics_points_earned, paper_points_earned
)
UPDATE accounts 
SET 
    bel_point_balance = bel_point_balance + new_waste.total_points_earned,
    total_bel_point = total_bel_point + new_waste.total_points_earned,
    total_glass_waste = total_glass_waste + new_waste.glass_weighed,
    total_metal_waste = total_metal_waste + new_waste.metal_weighed,
    total_plastics_waste = total_plastics_waste + new_waste.plastics_weighed,
    total_paper_waste = total_paper_waste + new_waste.paper_weighed,
    total_glass_points_earned = total_glass_points_earned + new_waste.glass_points_earned,
    total_metal_points_earned = total_metal_points_earned + new_waste.metal_points_earned,
    total_plastics_points_earned = total_plastics_points_earned + new_waste.plastics_points_earned,
    total_paper_points_earned = total_paper_points_earned + new_waste.paper_points_earned
FROM new_waste 
WHERE accounts.account_email = new_waste.account_email;

-- name: UpdateCollectorStats :exec
UPDATE collectors
SET 
    total_glass_waste = total_glass_waste + $1,
    total_metal_waste = total_metal_waste + $2,
    total_plastics_waste = total_plastics_waste + $3,
    total_paper_waste = total_paper_waste + $4
WHERE agent_ID = $5;


-- name: SoftDeleteCollector :exec
UPDATE agent
SET is_deleted = TRUE, deleted_at = NOW()
WHERE agent_ID = $1 AND agent_type = 'collector' AND is_deleted = FALSE;


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


-- name: GetAccountByEmail :one
SELECT * 
FROM accounts 
WHERE account_email = $1 
  AND is_Deleted = FALSE;


-- name: UpdateNickName :one
UPDATE accounts
SET display_name = $1
WHERE account_email = $2
RETURNING *;


-- name: GetAllCollectors :many
SELECT * FROM collectors;


-- name: GetCollector :one
SELECT * FROM collectors
WHERE collector_Email = $1;


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


-- name: GetWasteLogByAccountPaginated :many
SELECT * 
FROM waste_logs
WHERE account_email = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: GetWasteLogByCollectorPaginated :many
SELECT * 
FROM waste_logs
WHERE agent_ID = $1 
AND collected_route = 'collector'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: GetWasteByAccountForMonth :many
SELECT * 
FROM waste_logs
WHERE account_email = $1
AND created_at >= make_timestamp($2, $3, 1, 0, 0, 0) 
AND created_at < (make_timestamp($2, $3, 1, 0, 0, 0) + INTERVAL '1 month')
ORDER BY created_at DESC;


-- name: GetWasteByCollectorForMonth :one
WITH monthly_waste AS (
    SELECT 
        created_at, 
        glass_weighed, 
        metal_weighed, 
        plastics_weighed, 
        paper_weighed
    FROM waste_logs
    WHERE agent_ID = $1
    AND collected_route = 'collector'
    AND created_at >= make_timestamp($2, $3, 1, 0, 0, 0)
    AND created_at < (make_timestamp($2, $3, 1, 0, 0, 0) + INTERVAL '1 month')
    ORDER BY created_at ASC -- Ensure ordering is correct for the CTE itself
)
SELECT 
    COALESCE(SUM(glass_weighed)::DOUBLE PRECISION, 0) AS total_glass,
    COALESCE(SUM(metal_weighed)::DOUBLE PRECISION, 0) AS total_metal,
    COALESCE(SUM(plastics_weighed)::DOUBLE PRECISION, 0) AS total_plastics,
    COALESCE(SUM(paper_weighed)::DOUBLE PRECISION, 0) AS total_paper,

    COUNT(*) AS total_waste_logs,
    
    (
        SELECT JSON_AGG(glass_weighed ORDER BY created_at ASC)
        FROM monthly_waste
    ) AS glass_trend,
    (
        SELECT JSON_AGG(metal_weighed ORDER BY created_at ASC)
        FROM monthly_waste
    ) AS metal_trend,
    (
        SELECT JSON_AGG(plastics_weighed ORDER BY created_at ASC)
        FROM monthly_waste
    ) AS plastic_trend,
    (
        SELECT JSON_AGG(paper_weighed ORDER BY created_at ASC)
        FROM monthly_waste
    ) AS paper_trend
FROM monthly_waste;


-- name: GetAllNotifications :many
SELECT * FROM notifications ORDER BY notif_time DESC;


-- name: DeleteNotificationsByIDs :exec
DELETE FROM notifications WHERE id = ANY($1::int[]);
