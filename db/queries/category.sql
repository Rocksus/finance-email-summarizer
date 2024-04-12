-- name: InsertCategory :one
INSERT INTO category (
    category_name,
    created_at,
    updated_at
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateCategoryByID :exec
UPDATE category set category_name=$1, updated_at=NOW()
WHERE category_id=$2;

-- name: ListCategories :many
SELECT 
    category_id,
    category_name,
    parent_category_id,
    created_at,
    updated_at
FROM category;