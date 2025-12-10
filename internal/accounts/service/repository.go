package accounts

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountsRepository interface {
	GetAccount(id uuid.UUID) (Account, error)
	GetAllAccounts() ([]Account, error)
	CreateAccount(acc Account) (Account, error)
	UpdateAccount(acc Account) (Account, error)
	DeleteAccount(id uuid.UUID) error
}

type accountsRepository struct {
	db *gorm.DB
}

func NewAccountsRepository(db *gorm.DB) AccountsRepository {
	return &accountsRepository{db: db}
}

func (r *accountsRepository) GetAccount(id uuid.UUID) (Account, error) {
	var acc Account
	err := r.db.First(&acc, "id = ?", id).Error
	return acc, err
}

func (r *accountsRepository) GetAllAccounts() ([]Account, error) {
	var accounts []Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}

func (r *accountsRepository) CreateAccount(acc Account) (Account, error) {
	log.Printf("REPO: create account %s", acc.Name)
	return acc, r.db.Create(&acc).Error
}

func (r *accountsRepository) UpdateAccount(acc Account) (Account, error) {
	log.Printf("REPO: update account %s", acc.Name)
	return acc, r.db.Save(&acc).Error
}

func (r *accountsRepository) DeleteAccount(id uuid.UUID) error {
	log.Printf("REPO: delete account %s", id)
	return r.db.Delete(&Account{}, "id = ?", id).Error
}
