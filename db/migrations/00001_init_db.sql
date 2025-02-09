-- +goose Up
CREATE TABLE users (
    user_id    INTEGER PRIMARY KEY AUTOINCREMENT,
    name       VARCHAR(255) NOT NULL,
    username   VARCHAR(20) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_user_username ON users(username);

CREATE TABLE user_auth (
    user_id       INTEGER PRIMARY KEY,
    password_hash TEXT,
    google_id     TEXT,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE TABLE user_account (
    account_id                INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id                   INTEGER NOT NULL,
    account_name              VARCHAR(255) NOT NULL,
    magnified_balance_summary INTEGER NOT NULL,
    created_at                DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at                DATETIME,
    currency                  VARCHAR(3) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE INDEX idx_user_account_user_id ON user_account(user_id);

CREATE TABLE transaction_category (
    category_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name        VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  DATETIME
);

CREATE TABLE account_transaction (
    transaction_id                TEXT PRIMARY KEY,  -- UUID generated in app code
    account_id                    INTEGER NOT NULL,
    user_id                       INTEGER NOT NULL,
    transaction_name              VARCHAR(255) NOT NULL,
    magnified_transaction_amount  INTEGER NOT NULL,
    transaction_type              INTEGER NOT NULL,  -- 0 debit, 1 credit, 2 transfer
    category_id                   INTEGER NOT NULL,
    created_at                    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    notes                         VARCHAR(255) NOT NULL DEFAULT '',
    transaction_source            INTEGER,
    transaction_source_id         INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(user_id),
    FOREIGN KEY(account_id) REFERENCES user_account(account_id),
    FOREIGN KEY(category_id) REFERENCES transaction_category(category_id)
);

CREATE INDEX idx_account_transaction_account_id ON account_transaction(account_id);
CREATE INDEX idx_account_transaction_user_id ON account_transaction(user_id);
CREATE INDEX idx_account_transaction_category_id ON account_transaction(category_id);

CREATE TABLE account_subscription (
    subscription_id               INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id                    INTEGER NOT NULL,
    user_id                       INTEGER NOT NULL,
    bill_at                       DATETIME NOT NULL,
    transaction_type              INTEGER NOT NULL,  -- 0 debit, 1 credit, 2 transfer
    category_id                   INTEGER NOT NULL,
    notes                         VARCHAR(255) NOT NULL DEFAULT '',
    magnified_transaction_amount  INTEGER NOT NULL,
    created_at                    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at                    DATETIME,
    FOREIGN KEY(user_id) REFERENCES users(user_id),
    FOREIGN KEY(account_id) REFERENCES user_account(account_id),
    FOREIGN KEY(category_id) REFERENCES transaction_category(category_id)
);

CREATE INDEX idx_account_subscription_account_id ON account_subscription(account_id);

CREATE TABLE transaction_installment (
    installment_id                INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id                    INTEGER NOT NULL,
    user_id                       INTEGER NOT NULL,
    category_id                   INTEGER NOT NULL,
    transaction_type              INTEGER NOT NULL,  -- 0 debit, 1 credit, 2 transfer
    notes                         VARCHAR(255) NOT NULL DEFAULT '',
    magnified_transaction_amount  INTEGER NOT NULL,
    installment_plan_months       INTEGER NOT NULL,
    created_at                    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at                    DATETIME,
    FOREIGN KEY(user_id) REFERENCES users(user_id),
    FOREIGN KEY(account_id) REFERENCES user_account(account_id),
    FOREIGN KEY(category_id) REFERENCES transaction_category(category_id)
);

CREATE INDEX idx_transaction_installment_account_id ON transaction_installment(account_id);

CREATE TABLE currency_magnifier (
    currency  VARCHAR(3) PRIMARY KEY,
    magnifier INTEGER NOT NULL DEFAULT 1
);

-- +goose Down
DROP TABLE IF EXISTS currency_magnifier;
DROP TABLE IF EXISTS transaction_installment;
DROP TABLE IF EXISTS account_subscription;
DROP TABLE IF EXISTS account_transaction;
DROP TABLE IF EXISTS transaction_category;
DROP TABLE IF EXISTS user_account;
DROP TABLE IF EXISTS user_auth;
DROP TABLE IF EXISTS users;