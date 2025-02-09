-- ListCategories returns all non-deleted categories.
-- name: ListCategories :many
SELECT 
  category_id, 
  name, 
  description, 
  created_at, 
  updated_at
FROM transaction_category
WHERE deleted_at IS NULL
ORDER BY category_id;

-- CreateCategory inserts a new transaction category.
-- name: CreateCategory :one
INSERT INTO transaction_category (name, description)
VALUES ($1, $2)
RETURNING category_id, name, description, created_at, updated_at;

-- SoftDeleteCategory marks a category as deleted.
-- name: SoftDeleteCategory :one
UPDATE transaction_category
SET deleted_at = datetime('now')
WHERE category_id = $1
  AND deleted_at IS NULL
RETURNING category_id, name, description, created_at, updated_at;