package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type ComplaintRepository struct {
	db *gorm.DB
}

func NewComplaintRepository() *ComplaintRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &ComplaintRepository{db: db}
}

func (cr *ComplaintRepository) FindByID(id uuid.UUID) (*models.Complaint, error) {
	var m models.Complaint
	if err := cr.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (cr *ComplaintRepository) FindAll() ([]models.Complaint, error) {
	var complaints []models.Complaint
	if err := cr.db.Find(&complaints).Error; err != nil {
		return nil, err
	}
	return complaints, nil
}

func (cr *ComplaintRepository) Create(m *models.Complaint) error {
	return cr.db.Create(m).Error
}

func (cr *ComplaintRepository) Update(m *models.Complaint) error {
	return cr.db.Save(m).Error
}

func (cr *ComplaintRepository) Delete(m *models.Complaint) error {
	return cr.db.Delete(m).Error
}

func (cr *ComplaintRepository) DeleteByID(id uuid.UUID) bool {
	err := cr.db.Delete(&models.Complaint{}, id)
	if err != nil {
		return true
	}
	return false
}
