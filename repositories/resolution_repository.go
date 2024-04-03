package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type ResolutionRepository struct {
	db *gorm.DB
}

func NewResolutionRepository() *ResolutionRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &ResolutionRepository{db: db}
}

func (rr *ResolutionRepository) FindByID(id uuid.UUID) (*models.Resolution, error) {
	var m models.Resolution
	if err := rr.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (rr *ResolutionRepository) FindAll() ([]models.Resolution, error) {
	var resolutions []models.Resolution
	if err := rr.db.Find(&resolutions).Error; err != nil {
		return nil, err
	}
	return resolutions, nil
}

func (rr *ResolutionRepository) Create(m *models.Resolution) error {
	return rr.db.Create(m).Error
}

func (rr *ResolutionRepository) Update(m *models.Resolution) error {
	return rr.db.Save(m).Error
}

func (rr *ResolutionRepository) Delete(m *models.Resolution) error {
	return rr.db.Delete(m).Error
}

func (rr *ResolutionRepository) DeleteByID(id uuid.UUID) bool {
	err := rr.db.Delete(&models.Resolution{}, id)
	if err != nil {
		return true
	}
	return false
}
