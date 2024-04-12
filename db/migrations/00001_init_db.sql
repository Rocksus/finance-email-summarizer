-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_data (
    user_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(20) NOT NULL,
    email VARCHAR(20) NOT NULL,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL default CURRENT_TIMESTAMP
);

CREATE TABLE user_auth (
    user_id BIGINT PRIMARY KEY,
    password_hash text,
    FOREIGN KEY(user_id) REFERENCES user_data(user_id)
);

CREATE TABLE category (
    category_id BIGSERIAL PRIMARY KEY,
    category_name varchar(255) NOT NULL default '',
    parent_category_id BIGINT,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    FOREIGN KEY(parent_category_id) REFERENCES category(category_id)
);

CREATE TABLE user_account (
    account_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    magnified_balance_summary bigint NOT NULL,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    rollup_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    currency varchar(3) NOT NULL, -- ISO4217 coded currency
    FOREIGN KEY(user_id) REFERENCES user_data(user_id)
);

CREATE TABLE account_transaction (
    transaction_id UUID PRIMARY KEY,
    account_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    transaction_name varchar(255) NOT NULL,
    magnified_transaction_amount bigint NOT NULL,
    transaction_type int NOT NULL, -- 0 debit, 1 credit
    category_id BIGINT NOT NULL,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    notes varchar(255) NOT NULL default '',
    created_by varchar(255) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES user_data(user_id),
    FOREIGN KEY(account_id) REFERENCES user_account(account_id),
    FOREIGN KEY(category_id) REFERENCES category(category_id)
);

CREATE TABLE account_transaction_approval (
    approval_id UUID PRIMARY KEY,
    account_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    transaction_name varchar(255) NOT NULL,
    magnified_transaction_amount bigint NOT NULL,
    transaction_type int NOT NULL, -- 0 debit, 1 credit
    category_id BIGINT NOT NULL,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    notes varchar(255) NOT NULL default '',
    created_by varchar(255) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES user_data(user_id),
    FOREIGN KEY(account_id) REFERENCES user_account(account_id),
    FOREIGN KEY(category_id) REFERENCES category(category_id)
);


CREATE TABLE user_api_key (
    api_key_id VARCHAR(64) PRIMARY KEY,
    api_secret_hash VARCHAR(64) NOT NULL,
    identifier VARCHAR(255) NOT NULL,
    created_at timestamptz NOT NULL default CURRENT_TIMESTAMP,
    deleted_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE currency_magnifier;
DROP TABLE account_transaction;
DROP TABLE account_transaction_approval;
DROP TABLE user_api_key;
DROP TABLE user_account;
DROP TABLE user_auth;
DROP TABLE user_data;
DROP TABLE category_id;
-- +goose StatementEnd
