package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository() (*OrganizationRepository, error) {
	db, err := cmd.GetDB()
	if err != nil {
		return nil, err
	}
	return &OrganizationRepository{db: db}, nil
}

func (or *OrganizationRepository) FindByUser(user *models.User) ([]models.Organization, error) {
	var organization []models.Organization
	if err := or.db.Where("user_id = ?", user.ID).Find(&organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}

func (or *OrganizationRepository) FindByID(id uuid.UUID) (*models.Organization, error) {
	var Organization models.Organization
	if err := or.db.First(&Organization, id).Error; err != nil {
		return nil, err
	}
	return &Organization, nil
}

func (or *OrganizationRepository) FindAll() ([]models.Organization, error) {
	var Organizations []models.Organization
	if err := or.db.Find(&Organizations).Error; err != nil {
		return nil, err
	}
	return Organizations, nil
}

func (or *OrganizationRepository) Create(Organization *models.Organization) error {
	return or.db.Create(Organization).Error
}

func (or *OrganizationRepository) Update(Organization *models.Organization) error {
	return or.db.Save(Organization).Error
}

func (or *OrganizationRepository) Delete(Organization *models.Organization) error {
	return or.db.Delete(Organization).Error
}

func (or *OrganizationRepository) DeleteByID(id uuid.UUID) bool {
	err := or.db.Delete(&models.Organization{}, id)
	if err != nil {
		return true
	}
	return false
}
