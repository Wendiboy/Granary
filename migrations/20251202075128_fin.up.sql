CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE spends (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    type TEXT NOT NULL,

    account_id UUID NOT NULL,
    account_to_id UUID,

    category_id UUID NOT NULL,

    amount NUMERIC(12,2) NOT NULL,
    currency TEXT NOT NULL,

    labels TEXT[],              -- array of strings

    note TEXT,

    date TIMESTAMPTZ NOT NULL,

    is_pending BOOLEAN NOT NULL DEFAULT FALSE,

    metadata JSONB DEFAULT '{}'::jsonb,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
