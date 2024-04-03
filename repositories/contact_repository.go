package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository() *ContactRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &ContactRepository{db: db}
}

func (ur *ContactRepository) FindByUser(user *models.User) ([]models.Contact, error) {
	var Contacts []models.Contact
	if err := ur.db.Where("user_id = ?", user.ID).Find(&Contacts).Error; err != nil {
		return nil, err
	}
	return Contacts, nil
}

func (ur *ContactRepository) FindByUserID(id uuid.UUID) ([]models.Contact, error) {
	var Contacts []models.Contact
	if err := ur.db.Where("user_id = ?", id).Find(&Contacts).Error; err != nil {
		return nil, err
	}
	return Contacts, nil
}

func (ur *ContactRepository) FindByID(id uuid.UUID) (*models.Contact, error) {
	var Contact models.Contact
	if err := ur.db.First(&Contact, id).Error; err != nil {
		return nil, err
	}
	return &Contact, nil
}

func (ur *ContactRepository) FindDefaultByUser(user *models.User) (*models.Contact, error) {
	var Contact models.Contact
	if err := ur.db.Where("user_id = ? AND is_default = ?", user.ID, true).First(&Contact).Error; err != nil {
		return nil, err
	}
	return &Contact, nil
}

func (ur *ContactRepository) Create(Contact *models.Contact) error {
	return ur.db.Create(Contact).Error
}

func (ur *ContactRepository) Update(Contact *models.Contact) error {
	return ur.db.Save(Contact).Error
}

func (ur *ContactRepository) Delete(Contact *models.Contact) error {
	return ur.db.Delete(Contact).Error
}

func (ur *ContactRepository) FindAll() ([]models.Contact, error) {
	var Contacts []models.Contact
	if err := ur.db.Find(&Contacts).Error; err != nil {
		return nil, err
	}
	return Contacts, nil
}

func (ur *ContactRepository) DeleteByID(id uuid.UUID) bool {
	err := ur.db.Delete(&models.Contact{}, id)
	if err != nil {
		return true
	}
	return false
}
