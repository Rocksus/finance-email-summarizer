-- CreateUser inserts a new user.
-- name: CreateUser :one
INSERT INTO "user" (name, username)
VALUES ($1, $2)
RETURNING user_id, name, username, created_at, updated_at;

-- GetUserByID returns a user by ID.
-- name: GetUserByID :one
SELECT 
  user_id, 
  name, 
  username, 
  created_at, 
  updated_at
FROM "user"
WHERE user_id = $1;

-- GetUserByUsername returns a user by username.
-- name: GetUserByUsername :one
SELECT 
  user_id, 
  name, 
  username, 
  created_at, 
  updated_at
FROM "user"
WHERE username = $1;

-- ListUsers returns all users ordered by creation time.
-- name: ListUsers :many
SELECT 
  user_id, 
  name, 
  username, 
  created_at, 
  updated_at
FROM "user"
ORDER BY created_at DESC;