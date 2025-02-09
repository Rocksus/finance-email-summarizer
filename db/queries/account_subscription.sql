-- CreateAccountSubscription inserts a new subscription.
-- name: CreateAccountSubscription :one
INSERT INTO account_subscription (
    account_id, 
    user_id, 
    bill_at, 
    transaction_type, 
    category_id, 
    notes, 
    magnified_transaction_amount
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING 
    subscription_id, 
    account_id, 
    user_id, 
    bill_at, 
    transaction_type, 
    category_id, 
    notes, 
    magnified_transaction_amount, 
    created_at, 
    updated_at;

-- GetAccountSubscriptionByID returns a subscription by its ID.
-- name: GetAccountSubscriptionByID :one
SELECT 
    subscription_id, 
    account_id, 
    user_id, 
    bill_at, 
    transaction_type, 
    category_id, 
    notes, 
    magnified_transaction_amount, 
    created_at, 
    updated_at
FROM account_subscription
WHERE subscription_id = $1
  AND deleted_at IS NULL;

-- ListAccountSubscriptionsByAccount returns all non-deleted subscriptions for an account.
-- name: ListAccountSubscriptionsByAccount :many
SELECT 
    subscription_id, 
    account_id, 
    user_id, 
    bill_at, 
    transaction_type, 
    category_id, 
    notes, 
    magnified_transaction_amount, 
    created_at, 
    updated_at
FROM account_subscription
WHERE account_id = $1
  AND deleted_at IS NULL
ORDER BY bill_at DESC;

-- UpdateAccountSubscription updates an existing subscription.
-- name: UpdateAccountSubscription :one
UPDATE account_subscription
SET 
  bill_at = $2,
  transaction_type = $3,
  category_id = $4,
  notes = $5,
  magnified_transaction_amount = $6,
  updated_at = datetime('now')
WHERE subscription_id = $1
  AND deleted_at IS NULL
RETURNING 
    subscription_id, 
    account_id, 
    user_id, 
    bill_at, 
    transaction_type, 
    category_id, 
    notes, 
    magnified_transaction_amount, 
    created_at, 
    updated_at;