-- +goose Up
CREATE TABLE IF NOT EXISTS "accounts" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "full_name" varchar NOT NULL,
    "balance" money NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT balance_non_negative CHECK (balance::numeric::float >= 0)
);

CREATE TABLE IF NOT EXISTS "transactions" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "account_id" uuid REFERENCES "accounts" ("id"),
    "amount" money NOT NULL,
    "info" text NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "is_accrual" bool NOT NULL, -- whether to add or remove money from balance. If true, money are added to balance

    CONSTRAINT amount_non_negative CHECK (amount::numeric::float >= 0)
);

CREATE TABLE IF NOT EXISTS "reservations" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "account_id" uuid REFERENCES "accounts" ("id"),
    "service_id" uuid NOT NULL,
    "order_id" uuid NOT NULL,
    "amount" money NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "closed_at" timestamptz,
    "is_reservation_active" bool DEFAULT true,

    CONSTRAINT amount_non_negative CHECK (amount::numeric::float >= 0)
);

CREATE TABLE IF NOT EXISTS "reports" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "account_id" uuid NOT NULL REFERENCES "accounts" ("id"),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- +goose Down
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS reservations;
DROP TABLE IF EXISTS reports;