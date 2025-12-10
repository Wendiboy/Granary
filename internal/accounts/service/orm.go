package accounts

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	// Основные данные
	Name     string `gorm:"size:100;not null"` // "Основная карта", "Наличные", "Счёт в Т-Банке"
	BankName string `gorm:"size:100"`          // "Тинькофф", "Сбер", "Альфа-Банк", "Cash"
	Icon     string `gorm:"size:50"`           // "credit_card", "cash", "wallet", "piggy_bank" или путь к иконке
	Color    string `gorm:"size:7"`            // "#FF5733" (hex), для красивого отображения в UI

	// Финансовые характеристики
	Balance  float64 `gorm:"default:0"`       // текущий баланс (может быть отрицательным для кредиток)
	Currency string  `gorm:"size:3;not null"` // ISO 4217: "RUB", "USD", "EUR", "BTC"

	// Тип счёта — важная вещь для аналитики
	Type AccountType `gorm:"type:varchar(20);not null;index"` // enum в БД

	// Дополнительные поля для разных типов счетов
	InitialBalance float64    `gorm:"column:initial_balance;default:0"` // начальный баланс на момент открытия
	OpenedAt       time.Time  `gorm:"column:opened_at"`                 // дата открытия
	ClosedAt       *time.Time `gorm:"column:closed_at"`                 // дата закрытия (nil = активный)

	// Заметки и порядок
	Note      string
	SortOrder int `gorm:"default:0"` // для сортировки в списке счетов

	// Флаги
	IsActive bool `gorm:"default:true"`
	IsHidden bool `gorm:"default:false"` // скрыть из основного списка (например, старые закрытые счета)

	// GORM timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// DTO для ответа (GET /accounts, GET /accounts/:id)
type AccountResponseDTO struct {
	ID     string `json:"id"`      // uuid as string
	UserID string `json:"user_id"` // если нужно, иначе можно опустить

	Name     string `json:"name"`
	BankName string `json:"bank_name,omitempty"`
	Icon     string `json:"icon,omitempty"`  // "credit_card", "cash", "wallet" и т.д.
	Color    string `json:"color,omitempty"` // "#FF5733"

	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"` // "RUB", "USD", "BTC"

	Type string `json:"type"` // "cash", "debit_card", "credit_card", "savings" и т.д.

	InitialBalance float64 `json:"initial_balance"`
	OpenedAt       string  `json:"opened_at"` // ISO 8601: "2025-12-10T00:00:00Z"
	ClosedAt       *string `json:"closed_at,omitempty"`

	Note      string `json:"note,omitempty"`
	SortOrder int    `json:"sort_order"`

	IsActive bool `json:"is_active"`
	IsHidden bool `json:"is_hidden"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// DTO для создания нового счёта (POST /accounts)
type AccountCreateDTO struct {
	Name     string `json:"name" validate:"required,max=100"`
	BankName string `json:"bank_name,omitempty" validate:"max=100"`
	Icon     string `json:"icon,omitempty" validate:"max=50"`
	Color    string `json:"color,omitempty" validate:"omitempty,hexcolor"` // "#RRGGBB"

	Currency string `json:"currency" validate:"required,oneof=RUB USD EUR BTC ETH"` // можно расширить
	Type     string `json:"type" validate:"required,oneof=cash debit_card credit_card savings investment deposit loan crypto other"`

	InitialBalance string `json:"initial_balance" validate:"required"`     // обычно = 0, но можно задать
	OpenedAt       string `json:"opened_at" validate:"omitempty,datetime"` // ISO 8601, если пусто — берём now()
	ClosedAt       string `json:"closed_at,omitempty"`

	Note      string `json:"note,omitempty" validate:"max=500"`
	SortOrder string `json:"sort_order" validate:"omitempty,min=0"`

	IsActive bool `json:"is_active" validate:"omitempty"` // по умолчанию true
	IsHidden bool `json:"is_hidden" validate:"omitempty"` // по умолчанию false
}

// DTO для обновления (PATCH /accounts/:id)
type AccountUpdateDTO struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,max=100"`
	BankName *string `json:"bank_name,omitempty" validate:"omitempty,max=100"`
	Icon     *string `json:"icon,omitempty" validate:"omitempty,max=50"`
	Color    *string `json:"color,omitempty" validate:"omitempty,hexcolor"`

	Currency *string `json:"currency,omitempty" validate:"omitempty,oneof=RUB USD EUR BTC ETH"`
	Type     *string `json:"type,omitempty" validate:"omitempty,oneof=cash debit_card credit_card savings investment deposit loan crypto other"`

	InitialBalance *float64 `json:"initial_balance,omitempty"`
	OpenedAt       *string  `json:"opened_at,omitempty" validate:"omitempty,datetime"`
	ClosedAt       *string  `json:"closed_at,omitempty" validate:"omitempty,datetime"`

	Note      *string `json:"note,omitempty" validate:"omitempty,max=500"`
	SortOrder *int    `json:"sort_order,omitempty" validate:"omitempty,min=0"`

	IsActive *bool `json:"is_active,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
}

type AccountType string

const (
	AccountTypeCash       AccountType = "cash"        // наличные
	AccountTypeDebitCard  AccountType = "debit_card"  // дебетовая карта
	AccountTypeCreditCard AccountType = "credit_card" // кредитная карта
	AccountTypeSavings    AccountType = "savings"     // накопительный счёт
	AccountTypeInvestment AccountType = "investment"  // брокерский/инвестиционный
	AccountTypeDeposit    AccountType = "deposit"     // вклад/депозит
	AccountTypeLoan       AccountType = "loan"        // кредит/заём
	AccountTypeCrypto     AccountType = "crypto"      // криптокошелёк
	AccountTypeOther      AccountType = "other"       // прочее
)
