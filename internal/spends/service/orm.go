package SpendsService

import (
	"time"

	"github.com/google/uuid"
)

type Spend struct {
	ID       uuid.UUID `json:"id"`
	Account  string    `json:"account"`
	Category string    `json:"category"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	Labels   string    `json:"labels"`
	Note     string    `json:"note"`
	Date     string    `json:"date"`
}

type RawSpend struct {
	Id        uuid.UUID
	Account   string
	Category  string
	Amount    float64
	Currency  string
	Labels    string
	Note      string
	Date      string
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

// сделать тип для каждого поля (строгая типизация) enum
func MappingSpend(rawSpend RawSpend) Spend {
	spend := Spend{
		ID:       rawSpend.Id,
		Account:  rawSpend.Account,
		Category: rawSpend.Category,
		Amount:   rawSpend.Amount,
		Currency: rawSpend.Currency,
		Labels:   rawSpend.Labels,
		Note:     rawSpend.Note,
		Date:     rawSpend.Date,
	}
	return spend
}
