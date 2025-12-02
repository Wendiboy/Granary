CREATE TABLE raw_spends (
    id          UUID PRIMARY KEY,
    account     VARCHAR(255) NOT NULL,
    category    VARCHAR(255) NOT NULL,
    amount      NUMERIC(12, 2) NOT NULL,
    currency    VARCHAR(10) NOT NULL,
    labels      TEXT,
    note        TEXT,
    date        DATE NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
