package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/data/models"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository() *FileRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &FileRepository{db: db}
}

func (fr *FileRepository) FindByID(id uuid.UUID) (*models.File, error) {
	var f models.File
	if err := fr.db.First(&f, id).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

func (fr *FileRepository) FindAll() ([]models.File, error) {
	var files []models.File
	if err := fr.db.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (fr *FileRepository) Create(m *models.File) error {
	return fr.db.Create(m).Error
}

func (fr *FileRepository) Update(m *models.File) error {
	return fr.db.Save(m).Error
}

func (fr *FileRepository) Delete(m *models.File) error {
	return fr.db.Delete(m).Error
}

func (fr *FileRepository) DeleteByID(id uuid.UUID) bool {
	err := fr.db.Delete(&models.File{}, id)
	if err != nil {
		return true
	}
	return false
}
