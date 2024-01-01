CREATE TABLE user (
    user_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(20) NOT NULL,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL default CURRENT_TIMESTAMP
);

CREATE TABLE user_auth (
    user_id BIGINT PRIMARY KEY,
    password_hash text,
    google_id text,
    FOREIGN KEY(user_id) REFERENCES user(user_id)
);

CREATE TABLE user_account (
    account_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    magnified_balance_summary bigint NOT NULL,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    currency varchar(3) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES user(user_id)
);

CREATE TABLE account_transaction (
    transaction_id UUID PRIMARY KEY,
    account_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    transaction_name varchar(255) NOT NULL,
    magnified_transaction_amount bigint NOT NULL,
    transaction_type int NOT NULL, -- 0 debit, 1 credit
    category_id int NOT NULL,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    notes varchar(255) NOT NULL default '',
    FOREIGN KEY(user_id) REFERENCES user(user_id),
    FOREIGN KEY(account_id) REFERENCES user_account(account_id)
);

CREATE TABLE currency_magnifier (
    currency varchar(3) PRIMARY KEY,
    magnifier int  NOT NULL DEFAULT 1
);