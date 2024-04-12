-- name: InsertUserData :one
INSERT INTO user_data (
    name,
    username,
    email
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: InsertUserAuth :exec
INSERT INTO user_auth (
    user_id,
    password_hash
) VALUES (
    $1,
    $2
);

-- name: InsertUserAPIKey :one
INSERT INTO user_api_key (
    api_key_id,
    api_secret_hash,
    identifier,
    created_at
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: DeleteUserAPIKey :exec
UPDATE user_api_key
SET deleted_at = CURRENT_TIMESTAMP
WHERE api_key_id = $1;

-- name: GetUserAPIKey :one
SELECT 
    api_key_id,
    api_secret_hash,
    identifier,
    created_at,
    deleted_at
FROM user_api_key
WHERE api_key_id = $1;