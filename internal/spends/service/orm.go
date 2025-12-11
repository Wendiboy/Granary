package SpendsService

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type SpendRequestDTO struct {
	Id                    string `json:"id"`
	LinkedTransferSpendID string `json:"linked_spend_id"`

	Date        string `json:"date"`
	Type        string `json:"type"`
	AreaID      string `json:"area_id,omitempty"`
	CategoryID  string `json:"category_id"`
	AccountID   string `json:"account_id"`
	AccountToID string `json:"account_to_id,omitempty"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Note        string `json:"note,omitempty"`
	Labels      string `json:"labels"`
}

type Spend struct {
	ID          uuid.UUID
	Type        string
	AccountID   uuid.UUID
	AccountToID *uuid.UUID
	// AreaID      uuid.UUID
	CategoryID uuid.UUID
	Amount     float64
	Currency   string
	Labels     pq.StringArray `gorm:"type:text[]"`
	Note       string
	Date       time.Time

	IsPending bool
	// LinkedTransferSpendID *uuid.UUID `gorm:"type:uuid"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type SpendResponseDTO struct {
	ID                    string `json:"id"`
	LinkedTransferSpendID string `json:"linked_spend_id"`

	Date        string `json:"date"`
	Type        string `json:"type"`
	AreaID      string `json:"area_id"`
	CategoryID  string `json:"category_id"`
	AccountID   string `json:"account_id"`
	AccountToID string `json:"account_to_id,omitempty"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Note        string `json:"note"`
	Labels      string `json:"labels"`

	CreatedAt string `json:"created_at"`
}

func MappingSpend(req SpendRequestDTO) (Spend, error) {
	log.Println("ORM:", req)

	var err error

	// Parse ID (optional)
	var id uuid.UUID
	if req.Id != "" {
		id, err = uuid.Parse(req.Id)
		if err != nil {
			log.Println(fmt.Errorf("invalid id: %w", err))
			return Spend{}, fmt.Errorf("invalid id: %w", err)
		}
	} else {
		id = uuid.New()
	}

	// Parse AccountID
	accountID, err := uuid.Parse(req.AccountID)
	if err != nil {
		log.Println(fmt.Errorf("invalid account_id: %w", err))
		return Spend{}, fmt.Errorf("invalid account_id: %w", err)
	}

	// Parse AccountToID (optional)
	var accountToID *uuid.UUID
	if req.AccountToID != "" {
		parsed, err := uuid.Parse(req.AccountToID)
		if err != nil {
			log.Println(fmt.Errorf("invalid account_to_id: %w", err))
			return Spend{}, fmt.Errorf("invalid account_to_id: %w", err)
		}
		accountToID = &parsed
	}

	// Parse CategoryID
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		log.Println(fmt.Errorf("invalid account_to_id: %w", err))
		return Spend{}, fmt.Errorf("invalid category_id: %w", err)
	}

	// Parse Amount
	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		log.Println(fmt.Errorf("invalid amount: %w", err))
		return Spend{}, fmt.Errorf("invalid amount: %w", err)
	}

	// Parse Date (ISO-8601)
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		log.Println(fmt.Errorf("invalid account_to_id: %w", err))
		return Spend{}, fmt.Errorf("invalid date: %w", err)
	}

	// Parse Labels (string → []string)
	labels := []string{}
	if req.Labels != "" {
		labels = strings.Split(req.Labels, ",")
		for i := range labels {
			labels[i] = strings.TrimSpace(labels[i])
		}
	}

	// Construct domain model
	spend := Spend{
		ID:          id,
		Type:        req.Type,
		AccountID:   accountID,
		AccountToID: accountToID,
		CategoryID:  categoryID,
		Amount:      amount,
		Currency:    req.Currency,
		Labels:      labels,
		Note:        req.Note,
		Date:        date,

		// Default values
		IsPending: false,
	}
	log.Println("ORM2:", spend)
	return spend, nil
}

func MapSpendToResponseDTO(s Spend) SpendResponseDTO {
	// Преобразуем Labels []string → "a,b,c"
	labels := strings.Join(s.Labels, ",")

	return SpendResponseDTO{
		ID:         s.ID.String(),
		AccountID:  s.AccountID.String(),
		CategoryID: s.CategoryID.String(),
		Type:       s.Type,
		Amount:     fmt.Sprintf("%.2f", s.Amount), // float64 → "12.30"
		Currency:   s.Currency,
		Note:       s.Note,
		Labels:     labels,
		Date:       s.Date.Format(time.RFC3339),
		CreatedAt:  s.CreatedAt.Format(time.RFC3339),
	}
}
