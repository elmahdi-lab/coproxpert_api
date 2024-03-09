package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository() (*TokenRepository, error) {
	db, err := cmd.GetDB()
	if err != nil {
		return nil, err
	}
	return &TokenRepository{db: db}, nil
}

func (ur *TokenRepository) FindByID(id uuid.UUID) (*models.Token, error) {
	var Token models.Token
	if err := ur.db.First(&Token, id).Error; err != nil {
		return nil, err
	}
	return &Token, nil
}

func (ur *TokenRepository) Create(Token *models.Token) error {
	return ur.db.Create(Token).Error
}

func (ur *TokenRepository) Update(Token *models.Token) error {
	return ur.db.Save(Token).Error
}

func (ur *TokenRepository) Delete(Token *models.Token) error {
	return ur.db.Delete(Token).Error
}

func (ur *TokenRepository) FindAll() ([]models.Token, error) {
	var Tokens []models.Token
	if err := ur.db.Find(&Tokens).Error; err != nil {
		return nil, err
	}
	return Tokens, nil
}
