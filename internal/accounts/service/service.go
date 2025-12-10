package accounts

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type AccountsService interface {
	GetAccount(id uuid.UUID) (AccountResponseDTO, error)
	GetAllAccounts() ([]AccountResponseDTO, error)
	CreateAccount(req AccountCreateDTO) (AccountResponseDTO, error)
	UpdateAccount(id uuid.UUID, req AccountUpdateDTO) (AccountResponseDTO, error)
	DeleteAccount(id uuid.UUID) error
}

type accountsService struct {
	repo AccountsRepository
}

func NewAccountsService(repo AccountsRepository) AccountsService {
	return &accountsService{repo: repo}
}

// Маппинг из модели в DTO
func toResponseDTO(acc Account) AccountResponseDTO {
	closedAt := (*string)(nil)
	if acc.ClosedAt != nil {
		s := acc.ClosedAt.Format(time.RFC3339)
		closedAt = &s
	}

	return AccountResponseDTO{
		ID:             acc.ID.String(),
		Name:           acc.Name,
		BankName:       acc.BankName,
		Icon:           acc.Icon,
		Color:          acc.Color,
		Balance:        acc.Balance,
		Currency:       acc.Currency,
		Type:           string(acc.Type),
		InitialBalance: acc.InitialBalance,
		OpenedAt:       acc.OpenedAt.Format(time.RFC3339),
		ClosedAt:       closedAt,
		Note:           acc.Note,
		SortOrder:      acc.SortOrder,
		IsActive:       acc.IsActive,
		IsHidden:       acc.IsHidden,
		CreatedAt:      acc.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      acc.UpdatedAt.Format(time.RFC3339),
	}
}

// Маппинг из CreateDTO в модель
func fromCreateDTO(req AccountCreateDTO) (Account, error) {

	// --- parse balance
	initialBalance, err := strconv.ParseFloat(req.InitialBalance, 64)
	if err != nil {
		return Account{}, fmt.Errorf("invalid initial_balance: %w", err)
	}

	// --- parse sort order
	var sortOrder int
	if req.SortOrder != "" {
		sortOrder, err = strconv.Atoi(req.SortOrder)
		if err != nil {
			return Account{}, fmt.Errorf("invalid sort_order: %w", err)
		}
	}

	// --- parse opened_at
	var openedAt time.Time
	if req.OpenedAt == "" {
		openedAt = time.Now()
	} else {
		openedAt, err = time.Parse("2006-01-02", req.OpenedAt) // HTML <input type=date>
		if err != nil {
			return Account{}, fmt.Errorf("invalid opened_at: %w", err)
		}
	}

	// --- parse closed_at
	var closedAt *time.Time
	if req.ClosedAt != "" {
		t, err := time.Parse("2006-01-02", req.ClosedAt)
		if err != nil {
			return Account{}, fmt.Errorf("invalid closed_at: %w", err)
		}
		closedAt = &t
	}

	return Account{
		Name:           req.Name,
		BankName:       req.BankName,
		Icon:           req.Icon,
		Color:          req.Color,
		Currency:       req.Currency,
		Type:           AccountType(req.Type),
		InitialBalance: initialBalance,
		OpenedAt:       openedAt,
		ClosedAt:       closedAt,
		Note:           req.Note,
		SortOrder:      sortOrder,
		IsActive:       req.IsActive,
		IsHidden:       req.IsHidden,
	}, nil
}

// Маппинг из UpdateDTO в модель (только заполненные поля)
func fromUpdateDTO(acc Account, req AccountUpdateDTO) Account {
	if req.Name != nil {
		acc.Name = *req.Name
	}
	if req.BankName != nil {
		acc.BankName = *req.BankName
	}
	if req.Icon != nil {
		acc.Icon = *req.Icon
	}
	if req.Color != nil {
		acc.Color = *req.Color
	}
	if req.Currency != nil {
		acc.Currency = *req.Currency
	}
	if req.Type != nil {
		acc.Type = AccountType(*req.Type)
	}
	if req.InitialBalance != nil {
		acc.InitialBalance = *req.InitialBalance
	}
	if req.OpenedAt != nil {
		t, _ := time.Parse(time.RFC3339, *req.OpenedAt)
		acc.OpenedAt = t
	}
	if req.Note != nil {
		acc.Note = *req.Note
	}
	if req.SortOrder != nil {
		acc.SortOrder = *req.SortOrder
	}
	if req.IsActive != nil {
		acc.IsActive = *req.IsActive
	}
	if req.IsHidden != nil {
		acc.IsHidden = *req.IsHidden
	}
	return acc
}

func (s *accountsService) GetAccount(id uuid.UUID) (AccountResponseDTO, error) {
	acc, err := s.repo.GetAccount(id)
	if err != nil {
		return AccountResponseDTO{}, err
	}
	return toResponseDTO(acc), nil
}

func (s *accountsService) GetAllAccounts() ([]AccountResponseDTO, error) {
	accounts, err := s.repo.GetAllAccounts()
	if err != nil {
		return nil, err
	}

	dtos := make([]AccountResponseDTO, len(accounts))
	for i, acc := range accounts {
		dtos[i] = toResponseDTO(acc)
	}
	return dtos, nil
}

func (s *accountsService) CreateAccount(req AccountCreateDTO) (AccountResponseDTO, error) {
	acc, err := fromCreateDTO(req)
	if err != nil {
		return AccountResponseDTO{}, err
	}

	created, err := s.repo.CreateAccount(acc)
	if err != nil {
		return AccountResponseDTO{}, err
	}

	return toResponseDTO(created), nil
}

func (s *accountsService) UpdateAccount(id uuid.UUID, req AccountUpdateDTO) (AccountResponseDTO, error) {
	acc, err := s.repo.GetAccount(id)
	if err != nil {
		return AccountResponseDTO{}, err
	}

	updated := fromUpdateDTO(acc, req)

	updated, err = s.repo.UpdateAccount(updated)
	if err != nil {
		return AccountResponseDTO{}, err
	}

	return toResponseDTO(updated), nil
}

func (s *accountsService) DeleteAccount(id uuid.UUID) error {
	return s.repo.DeleteAccount(id)
}
