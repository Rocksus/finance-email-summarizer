-- CreateAccountTransaction inserts a new transaction.
-- name: CreateAccountTransaction :one
INSERT INTO account_transaction (
    transaction_id, 
    account_id, 
    user_id, 
    transaction_name, 
    magnified_transaction_amount, 
    transaction_type, 
    category_id, 
    notes, 
    transaction_source, 
    transaction_source_id
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING 
    transaction_id, 
    account_id, 
    user_id, 
    transaction_name, 
    magnified_transaction_amount, 
    transaction_type, 
    category_id, 
    created_at, 
    updated_at, 
    notes, 
    transaction_source, 
    transaction_source_id;

-- GetAccountTransactionByID returns a transaction by its ID.
-- name: GetAccountTransactionByID :one
SELECT 
    transaction_id, 
    account_id, 
    user_id, 
    transaction_name, 
    magnified_transaction_amount, 
    transaction_type, 
    category_id, 
    created_at, 
    updated_at, 
    notes, 
    transaction_source, 
    transaction_source_id
FROM account_transaction
WHERE transaction_id = $1;

-- ListAccountTransactionsByAccount returns all transactions for an account.
-- name: ListAccountTransactionsByAccount :many
SELECT 
    transaction_id, 
    account_id, 
    user_id, 
    transaction_name, 
    magnified_transaction_amount, 
    transaction_type, 
    category_id, 
    created_at, 
    updated_at, 
    notes, 
    transaction_source, 
    transaction_source_id
FROM account_transaction
WHERE account_id = $1
ORDER BY created_at DESC;