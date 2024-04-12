-- name: ListTransactionByAccountID :many
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
    created_by
FROM
    account_transaction
WHERE
    account_id = $1
LIMIT
    $2 OFFSET $3;

-- name: InsertAccountTransaction :one
INSERT INTO
    account_transaction (
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
        created_by
    ) VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11
    ) RETURNING *;

-- name: InsertUserAccount :one
INSERT INTO
    user_account (
        user_id,
        account_name,
        currency,
        magnified_balance_summary
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetUserAccount :one
SELECT
    ua.account_id,
    ua.user_id,
    ua.account_name,
    ua.magnified_balance_summary + COALESCE(
        (
            SELECT
                SUM(
                    CASE
                        WHEN at.transaction_type = 0 THEN at.magnified_transaction_amount
                        WHEN at.transaction_type = 1 THEN - at.magnified_transaction_amount
                    END
                )
            FROM
                account_transaction at
            WHERE
                at.account_id = ua.account_id
                AND at.created_at > ua.rollup_at
        ),
        0
    ) AS magnified_balance_summary,
    ua.created_at,
    ua.updated_at,
    ua.rollup_at,
    ua.currency
FROM
    user_account ua
WHERE
    ua.account_id = $1;

-- name: RollupUserAccountBalance :exec
UPDATE
    user_account ua
SET
    magnified_balance_summary = magnified_balance_summary + COALESCE(
        (
            SELECT
                SUM(
                    CASE
                        WHEN at.transaction_type = 0 THEN at.magnified_transaction_amount
                        WHEN at.transaction_type = 1 THEN - at.magnified_transaction_amount
                    END
                )
            FROM
                account_transaction at
            WHERE
                at.account_id = user_account.account_id
                AND at.created_at > user_account.rollup_at
                AND at.created_at <= $1
        ),
        0
    ),
    rollup_at = $1
WHERE
    ua.account_id = $2;