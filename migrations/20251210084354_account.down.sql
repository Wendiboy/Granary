-- Удаляем триггер
DROP TRIGGER IF EXISTS accounts_updated_at_trigger ON accounts;

-- Удаляем функцию
DROP FUNCTION IF EXISTS update_updated_at();

-- Удаляем check-констрейнт
ALTER TABLE accounts DROP CONSTRAINT IF EXISTS accounts_color_check;

-- Удаляем индексы
DROP INDEX IF EXISTS idx_accounts_type;
DROP INDEX IF EXISTS idx_accounts_opened_at;
DROP INDEX IF EXISTS idx_accounts_closed_at;
DROP INDEX IF EXISTS idx_accounts_deleted_at;

-- Удаляем саму таблицу
DROP TABLE IF EXISTS accounts;

-- Удаляем тип ENUM (осторожно — только если он больше нигде не используется!)
DROP TYPE IF EXISTS account_type CASCADE;