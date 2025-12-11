
type Label struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string         `gorm:"size:100;not null;uniqueIndex:idx_labels_user_name"` // "командировка", "подарок", "шаурма"
	Color       string         `gorm:"size:7"`   // "#FF6B6B" — для красивого отображения

	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type LabelResponseDTO struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Color       string `json:"color,omitempty"`
}

---

```go
type Label struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID      uuid.UUID      `gorm:"type:uuid;index"` // мультипользовательское приложение

	// Основные данные
	Name        string         `gorm:"size:100;not null;uniqueIndex:idx_labels_user_name"` // "командировка", "подарок", "шаурма"
	Description string         `gorm:"size:255"` // необязательное описание
	Color       string         `gorm:"size:7"`   // "#FF6B6B" — для красивого отображения
	Icon        string         `gorm:"size:50"`  // "flight", "gift", "emoji:🎁"

	// Флаги
	IsActive    bool           `gorm:"default:true"`   // можно архивировать
	IsSystem    bool           `gorm:"default:false"`  // системные теги (например, "transfer")
	SortOrder   int            `gorm:"default:0"`      // порядок в списке

	// GORM timestamps + soft-delete
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
```

```go
type LabelResponseDTO struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    Color       string `json:"color,omitempty"`
    Icon        string `json:"icon,omitempty"`
    IsActive    bool   `json:"is_active"`
    IsSystem    bool   `json:"is_system"`
    SortOrder   int    `json:"sort_order"`
}
```
```sql
CREATE TABLE IF NOT EXISTS labels (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID NOT NULL,

    name        VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    color       VARCHAR(7),
    icon        VARCHAR(50),

    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    is_system   BOOLEAN NOT NULL DEFAULT FALSE,
    sort_order  INTEGER NOT NULL DEFAULT 0,

    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,

    -- Уникальность: один пользователь — один тег с таким именем
    CONSTRAINT labels_user_name_unique UNIQUE (user_id, name)
);

CREATE INDEX idx_labels_user_id     ON labels(user_id);
CREATE INDEX idx_labels_is_active   ON labels(is_active);
CREATE INDEX idx_labels_deleted_at  ON labels(deleted_at);

-- Триггер updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER labels_updated_at_trigger
    BEFORE UPDATE ON labels
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();
```

```sql
DROP TRIGGER IF EXISTS labels_updated_at_trigger ON labels;
DROP FUNCTION IF EXISTS update_updated_at();

DROP INDEX IF EXISTS idx_labels_user_id;
DROP INDEX IF EXISTS idx_labels_is_active;
DROP INDEX IF EXISTS idx_labels_deleted_at;

DROP TABLE IF EXISTS labels;
```