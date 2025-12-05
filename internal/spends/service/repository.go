package SpendsService

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpendsRepository interface {
	GetSpend(id uuid.UUID) (Spend, error)
	GetAllSpends() ([]Spend, error)
	CreateSpend(rawSpend Spend) (Spend, error)
	UpdateSpend(rawSpend Spend) (Spend, error)
	DeleteSpend(id uuid.UUID) error
}

type spendsRepository struct {
	db *gorm.DB
}

func NewSpendsRepository(db *gorm.DB) SpendsRepository {
	return &spendsRepository{db: db}
}

func (r *spendsRepository) GetSpend(id uuid.UUID) (Spend, error) {
	var rawSpend Spend
	err := r.db.Find(&rawSpend, "id=?", id).Error
	return rawSpend, err
}

func (r *spendsRepository) GetAllSpends() ([]Spend, error) {
	log.Println("REPO: GetAllSpends")
	var rawSpends []Spend
	err := r.db.Find(&rawSpends).Error
	return rawSpends, err
}

func (r *spendsRepository) CreateSpend(rawSpend Spend) (Spend, error) {
	log.Println("REPO: create", rawSpend)
	return rawSpend, r.db.Create(&rawSpend).Error
}

func (r *spendsRepository) UpdateSpend(rawSpend Spend) (Spend, error) {
	log.Println("REPO: update", rawSpend)

	return rawSpend, r.db.Save(&rawSpend).Error
}

func (r *spendsRepository) DeleteSpend(id uuid.UUID) error {
	log.Println("REPO: delete", id)
	return r.db.Delete(&Spend{}, "id=?", id).Error
}
