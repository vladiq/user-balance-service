CREATE TABLE "accounts" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "balance" DECIMAL(15, 2) NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT balance_non_negative CHECK (balance >= 0)
);

CREATE TABLE "transfers" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "account_id" uuid REFERENCES "accounts" ("id"),
    "is_accrual" bool NOT NULL, -- whether to add or remove money from balance. If true, money is added to balance
    "amount" DECIMAL(15, 2) NOT NULL,
    "info" text NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT amount_non_negative CHECK (amount >= 0)
);

CREATE TABLE "reservations" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "account_id" uuid REFERENCES "accounts" ("id"),
    "service_id" uuid NOT NULL,
    "order_id" uuid NOT NULL,
    "amount" DECIMAL(15, 2) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT amount_non_negative CHECK (amount >= 0)
);

CREATE TABLE "reports" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
    "service_id" uuid,
    "amount" DECIMAL(15, 2) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT amount_non_negative CHECK (amount >= 0)
);