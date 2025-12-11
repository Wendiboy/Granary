
Category
    ID 
    Name
    Color
 
---- 20%
type Category struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string         `gorm:"size:100;not null"` // "Продукты", "Зарплата", "Перевод на инвестиции"
	Color       string         `gorm:"size:7"`            // "#FF6B6B"

	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

-------- 

``` go
const (
	CategoryTypeExpense CategoryType = "expense" // расход
	CategoryTypeIncome  CategoryType = "income"  // доход
	CategoryTypeTransfer CategoryType = "transfer" // переводы (не влияют на общий баланс)
)

type Category struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID      uuid.UUID      `gorm:"type:uuid;index"` // для мультипользовательских приложений

	// Основные данные
	Name        string         `gorm:"size:100;not null"` // "Продукты", "Зарплата", "Перевод на инвестиции"
	Icon        string         `gorm:"size:50"`           // "food", "salary", "transfer", "emoji:🍔"
	Color       string         `gorm:"size:7"`            // "#FF6B6B"

	// Иерархия (группы категорий)
	ParentID    *uuid.UUID     `gorm:"type:uuid"` // nil = основная категория, иначе — подкатегория
	Parent      *Category      `gorm:"foreignKey:ParentID"` // для GORM

	// Тип категории
	Type        CategoryType   `gorm:"type:varchar(20);not null;index"`

	// Дополнительные флаги
	IsActive    bool           `gorm:"default:true"` // можно скрывать старые категории
	IsSystem    bool           `gorm:"default:false"` // системные категории (нельзя удалить)
	SortOrder   int            `gorm:"default:0"` // порядок в списке

	// GORM timestamps + soft-delete
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
```
```go
type CategoryResponseDTO struct {
    ID        string  `json:"id"`
    Name      string  `json:"name"`
    Icon      string  `json:"icon,omitempty"`
    Color     string  `json:"color,omitempty"`
    Type      string  `json:"type"`
    ParentID  *string `json:"parent_id,omitempty"`
    IsActive  bool    `json:"is_active"`
    IsSystem  bool    `json:"is_system"`
    SortOrder int     `json:"sort_order"`
}
```
```sql
DROP TRIGGER IF EXISTS categories_updated_at_trigger ON categories;
DROP FUNCTION IF EXISTS update_updated_at();

DROP INDEX IF EXISTS idx_categories_user_id;
DROP INDEX IF EXISTS idx_categories_type;
DROP INDEX IF EXISTS idx_categories_parent_id;
DROP INDEX IF EXISTS idx_categories_deleted_at;

DROP TABLE IF EXISTS categories;
DROP TYPE IF EXISTS category_type CASCADE;
```


```sql
-- ENUM для типа категории
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'category_type') THEN
        CREATE TYPE category_type AS ENUM ('expense', 'income', 'transfer');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS categories (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID NOT NULL,

    name        VARCHAR(100) NOT NULL,
    icon        VARCHAR(50),
    color       VARCHAR(7),

    parent_id   UUID,
    type        category_type NOT NULL,

    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    is_system   BOOLEAN NOT NULL DEFAULT FALSE,
    sort_order  INTEGER NOT NULL DEFAULT 0,

    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,

    CONSTRAINT fk_parent FOREIGN KEY (parent_id) REFERENCES categories(id) ON DELETE SET NULL
);

CREATE INDEX idx_categories_user_id   ON categories(user_id);
CREATE INDEX idx_categories_type      ON categories(type);
CREATE INDEX idx_categories_parent_id ON categories(parent_id);
CREATE INDEX idx_categories_deleted_at ON categories(deleted_at);

-- Триггер updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER categories_updated_at_trigger
    BEFORE UPDATE ON categories
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
```