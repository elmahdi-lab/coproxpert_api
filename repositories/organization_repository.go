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

func (ur *OrganizationRepository) FindByUser(user *models.User) ([]models.Organization, error) {
	var organization []models.Organization
	if err := ur.db.Where("user_id = ?", user.ID).Find(&organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}

func (ur *OrganizationRepository) FindByID(id uuid.UUID) (*models.Organization, error) {
	var Organization models.Organization
	if err := ur.db.First(&Organization, id).Error; err != nil {
		return nil, err
	}
	return &Organization, nil
}

func (ur *OrganizationRepository) Create(Organization *models.Organization) error {
	return ur.db.Create(Organization).Error
}

func (ur *OrganizationRepository) Update(Organization *models.Organization) error {
	return ur.db.Save(Organization).Error
}

func (ur *OrganizationRepository) Delete(Organization *models.Organization) error {
	return ur.db.Delete(Organization).Error
}

func (ur *OrganizationRepository) FindAll() ([]models.Organization, error) {
	var Organizations []models.Organization
	if err := ur.db.Find(&Organizations).Error; err != nil {
		return nil, err
	}
	return Organizations, nil
}

func (ur *OrganizationRepository) DeleteByID(id uuid.UUID) bool {
	err := ur.db.Delete(&models.Organization{}, id)
	if err != nil {
		return true
	}
	return false
}
