-- Создаём тип ENUM для account_type
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'account_type') THEN
        CREATE TYPE account_type AS ENUM (
            'cash',
            'debit_card',
            'credit_card',
            'savings',
            'investment',
            'deposit',
            'loan',
            'crypto',
            'other'
        );
    END IF;
END $$;

-- Таблица accounts
CREATE TABLE IF NOT EXISTS accounts (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name            VARCHAR(100) NOT NULL,
    bank_name       VARCHAR(100),
    icon            VARCHAR(50),
    color           VARCHAR(7),                 -- #RRGGBB

    balance         DOUBLE PRECISION NOT NULL DEFAULT 0,
    currency        VARCHAR(3) NOT NULL,        -- ISO 4217

    type            account_type NOT NULL,

    initial_balance DOUBLE PRECISION NOT NULL DEFAULT 0,
    opened_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    closed_at       TIMESTAMPTZ,

    note            TEXT,
    sort_order      INTEGER NOT NULL DEFAULT 0,

    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    is_hidden       BOOLEAN NOT NULL DEFAULT FALSE,

    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

-- Индексы
CREATE INDEX IF NOT EXISTS idx_accounts_type         ON accounts(type);
CREATE INDEX IF NOT EXISTS idx_accounts_opened_at    ON accounts(opened_at);
CREATE INDEX IF NOT EXISTS idx_accounts_closed_at    ON accounts(closed_at);
CREATE INDEX IF NOT EXISTS idx_accounts_deleted_at   ON accounts(deleted_at);

-- Опционально: проверка на корректный hex-цвет
ALTER TABLE accounts
    ADD CONSTRAINT accounts_color_check
    CHECK (color IS NULL OR color ~ '^#[0-9A-Fa-f]{6}$');

-- Триггер для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER accounts_updated_at_trigger
    BEFORE UPDATE ON accounts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();