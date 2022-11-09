-- +goose Up
CREATE TABLE IF NOT EXISTS "accounts" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "balance" DECIMAL(15, 2) NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT balance_non_negative CHECK (balance >= 0)
);

CREATE TABLE IF NOT EXISTS "transactions" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "account_id" uuid REFERENCES "accounts" ("id"),
    "is_accrual" bool NOT NULL, -- whether to add or remove money from balance. If true, money is added to balance
    "amount" DECIMAL(15, 2) NOT NULL,
    "info" text NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT amount_non_negative CHECK (amount >= 0)
);

CREATE TABLE IF NOT EXISTS "reservations" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "account_id" uuid REFERENCES "accounts" ("id"),
    "service_id" uuid NOT NULL,
    "order_id" uuid NOT NULL,
    "amount" DECIMAL(15, 2) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "closed_at" timestamptz,

    CONSTRAINT amount_non_negative CHECK (amount >= 0)
);

CREATE TABLE IF NOT EXISTS "reports" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "account_id" uuid NOT NULL REFERENCES "accounts" ("id"),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- +goose Down
DROP TABLE IF EXISTS accounts CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS reservations CASCADE;
DROP TABLE IF EXISTS reports CASCADE;