-- CreateTransactionInstallment inserts a new installment transaction.
-- name: CreateTransactionInstallment :one
INSERT INTO transaction_installment (
    account_id, 
    user_id, 
    category_id, 
    transaction_type, 
    notes, 
    magnified_transaction_amount, 
    installment_plan_months
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING 
    installment_id, 
    account_id, 
    user_id, 
    category_id, 
    transaction_type, 
    notes, 
    magnified_transaction_amount, 
    installment_plan_months, 
    created_at, 
    updated_at;

-- GetTransactionInstallmentByID returns an installment by its ID.
-- name: GetTransactionInstallmentByID :one
SELECT 
    installment_id, 
    account_id, 
    user_id, 
    category_id, 
    transaction_type, 
    notes, 
    magnified_transaction_amount, 
    installment_plan_months, 
    created_at, 
    updated_at
FROM transaction_installment
WHERE installment_id = $1
  AND deleted_at IS NULL;

-- ListTransactionInstallmentsByAccount returns all non-deleted installments for an account.
-- name: ListTransactionInstallmentsByAccount :many
SELECT 
    installment_id, 
    account_id, 
    user_id, 
    category_id, 
    transaction_type, 
    notes, 
    magnified_transaction_amount, 
    installment_plan_months, 
    created_at, 
    updated_at
FROM transaction_installment
WHERE account_id = $1
  AND deleted_at IS NULL
ORDER BY created_at DESC;

-- UpdateTransactionInstallment updates an existing installment.
-- name: UpdateTransactionInstallment :one
UPDATE transaction_installment
SET 
  category_id = $2,
  transaction_type = $3,
  notes = $4,
  magnified_transaction_amount = $5,
  installment_plan_months = $6,
  updated_at = datetime('now')
WHERE installment_id = $1
  AND deleted_at IS NULL
RETURNING 
    installment_id, 
    account_id, 
    user_id, 
    category_id, 
    transaction_type, 
    notes, 
    magnified_transaction_amount, 
    installment_plan_months, 
    created_at, 
    updated_at;