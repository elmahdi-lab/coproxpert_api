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

func NewOrganizationRepository() *OrganizationRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &OrganizationRepository{db: db}
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

func (or *OrganizationRepository) FindByUnitID(unitID uuid.UUID) (*models.Organization, error) {
	var organization models.Organization

	query := or.db.
		Joins("JOIN unit_groups ON unit_groups.id = units.unit_group_id").
		Joins("JOIN organizations ON organizations.id = unit_groups.organization_id").
		Where("units.id = ?", unitID).
		First(&organization)

	if query.Error != nil {
		return nil, query.Error
	}

	return &organization, nil
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
