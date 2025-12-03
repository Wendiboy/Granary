package SpendsService

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Spend struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Category string `json:"category"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Labels   string `json:"labels"`
	Note     string `json:"note"`
	Date     string `json:"date"`
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
		ID:       rawSpend.Id.String(),
		Account:  rawSpend.Account,
		Category: rawSpend.Category,
		Amount:   strconv.FormatFloat(rawSpend.Amount, 'f', 2, 64),
		Currency: rawSpend.Currency,
		Labels:   rawSpend.Labels,
		Note:     rawSpend.Note,
		Date:     rawSpend.Date,
	}
	return spend
}

func ReMappingSpend(id uuid.UUID, spend Spend) (RawSpend, error) {
	amount, err := strconv.ParseFloat(spend.Amount, 64)

	if err != nil {
		return RawSpend{}, err
	}

	rawSpend := RawSpend{
		Id:       id,
		Account:  spend.Account,
		Category: spend.Category,
		Amount:   amount,
		Currency: spend.Currency,
		Labels:   spend.Labels,
		Note:     spend.Note,
		Date:     spend.Date,
	}

	return rawSpend, nil
}
