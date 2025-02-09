-- CreateUserAuth inserts a new record in user_auth.
-- name: CreateUserAuth :one
INSERT INTO user_auth (user_id, password_hash, google_id)
VALUES ($1, $2, $3)
RETURNING user_id, password_hash, google_id;

-- GetUserAuthByUserID returns authentication details for a user.
-- name: GetUserAuthByUserID :one
SELECT 
  user_id, 
  password_hash, 
  google_id
FROM user_auth
WHERE user_id = $1;

-- UpdateUserAuth updates a user’s authentication information.
-- name: UpdateUserAuth :one
UPDATE user_auth
SET 
  password_hash = COALESCE($2, password_hash),
  google_id     = COALESCE($3, google_id)
WHERE user_id = $1
RETURNING user_id, password_hash, google_id;