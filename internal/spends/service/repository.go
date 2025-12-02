package SpendsService

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpendsRepository interface {
	GetSpend(id uuid.UUID) (RawSpend, error)
	GetAllSpends() ([]RawSpend, error)
	CreateSpend(rawSpend RawSpend) (RawSpend, error)
	UpdateSpend(rawSpend RawSpend) (RawSpend, error)
	DeleteSpend(id uuid.UUID) error
}

type spendsRepository struct {
	db *gorm.DB
}

func NewSpendsRepository(db *gorm.DB) SpendsRepository {
	return &spendsRepository{db: db}
}

func (r *spendsRepository) GetSpend(id uuid.UUID) (RawSpend, error) {
	var rawSpend RawSpend
	err := r.db.Find(&rawSpend, "id=?", id).Error
	return rawSpend, err
}

func (r *spendsRepository) GetAllSpends() ([]RawSpend, error) {
	log.Println("REPO: GetAllSpends")
	var rawSpends []RawSpend
	err := r.db.Find(&rawSpends).Error
	return rawSpends, err
}

func (r *spendsRepository) CreateSpend(rawSpend RawSpend) (RawSpend, error) {
	log.SetPrefix("REPO: create")
	log.Println(rawSpend)
	return rawSpend, r.db.Create(&rawSpend).Error
}

func (r *spendsRepository) UpdateSpend(rawSpend RawSpend) (RawSpend, error) {
	log.SetPrefix("REPO: Update")
	log.Println(rawSpend)

	return rawSpend, r.db.Save(&rawSpend).Error
}

func (r *spendsRepository) DeleteSpend(id uuid.UUID) error {
	return r.db.Delete(&RawSpend{}, "id=?", id).Error
}
