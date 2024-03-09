package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type CredentialRepository struct {
	db *gorm.DB
}

func NewCredentialRepository() (*CredentialRepository, error) {
	db, err := cmd.GetDB()
	if err != nil {
		return nil, err
	}
	return &CredentialRepository{db: db}, nil
}

func (ur *CredentialRepository) FindByID(id uuid.UUID) (*models.Credential, error) {
	var Credential models.Credential
	if err := ur.db.First(&Credential, id).Error; err != nil {
		return nil, err
	}
	return &Credential, nil
}

func (ur *CredentialRepository) Create(Credential *models.Credential) error {
	return ur.db.Create(Credential).Error
}

func (ur *CredentialRepository) Update(Credential *models.Credential) error {
	return ur.db.Save(Credential).Error
}

func (ur *CredentialRepository) Delete(Credential *models.Credential) error {
	return ur.db.Delete(Credential).Error
}

func (ur *CredentialRepository) FindAll() ([]models.Credential, error) {
	var Credentials []models.Credential
	if err := ur.db.Find(&Credentials).Error; err != nil {
		return nil, err
	}
	return Credentials, nil
}
