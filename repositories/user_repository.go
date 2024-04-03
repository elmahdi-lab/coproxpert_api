package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := ur.db.Preload("Permissions").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByToken(token string) (*models.User, error) {
	var user models.User
	if err := ur.db.Preload("Permissions").Where("token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := ur.db.Preload("Permissions").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (ur *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := ur.db.Preload("Permissions").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Create(user *models.User) error {
	id := ur.db.Create(user)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (ur *UserRepository) Update(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) Delete(user *models.User) error {
	return ur.db.Delete(user).Error
}

func (ur *UserRepository) DeleteByID(id uuid.UUID) bool {
	err := ur.db.Delete(&models.User{}, id)
	if err != nil {
		return false
	}
	return true
}
