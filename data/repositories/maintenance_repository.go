package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/data/models"
)

type MaintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository() *MaintenanceRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &MaintenanceRepository{db: db}
}

func (mr *MaintenanceRepository) FindByUnitGroup(ug *models.UnitGroup) ([]models.Maintenance, error) {
	var maintenance []models.Maintenance
	if err := mr.db.Where("unit_group_id = ?", ug.ID).Find(&maintenance).Error; err != nil {
		return nil, err
	}
	return maintenance, nil
}

func (mr *MaintenanceRepository) FindByUnit(u *models.Unit) ([]models.Maintenance, error) {
	var maintenance []models.Maintenance
	if err := mr.db.Where("unit_id = ?", u.ID).Find(&maintenance).Error; err != nil {
		return nil, err
	}
	return maintenance, nil
}

func (mr *MaintenanceRepository) FindByID(id uuid.UUID) (*models.Maintenance, error) {
	var m models.Maintenance
	if err := mr.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (mr *MaintenanceRepository) FindAll() ([]models.Maintenance, error) {
	var maintenances []models.Maintenance
	if err := mr.db.Find(&maintenances).Error; err != nil {
		return nil, err
	}
	return maintenances, nil
}

func (mr *MaintenanceRepository) Create(m *models.Maintenance) error {
	return mr.db.Create(m).Error
}

func (mr *MaintenanceRepository) Update(m *models.Maintenance) error {
	return mr.db.Save(m).Error
}

func (mr *MaintenanceRepository) Delete(m *models.Maintenance) error {
	return mr.db.Delete(m).Error
}

func (mr *MaintenanceRepository) DeleteByID(id uuid.UUID) bool {
	err := mr.db.Delete(&models.Maintenance{}, id)
	if err != nil {
		return true
	}
	return false
}
