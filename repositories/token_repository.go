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

func (tr *TokenRepository) FindByID(id uuid.UUID) (*models.Token, error) {
	var Token models.Token
	if err := tr.db.First(&Token, id).Error; err != nil {
		return nil, err
	}
	return &Token, nil
}

func (tr *TokenRepository) FindByToken(t string, preload bool) (*models.Token, error) {
	var token models.Token
	query := tr.db.Where("token = ?", t)

	if preload {
		query = query.Preload("User")
	}

	if err := query.First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (tr *TokenRepository) Create(token *models.Token) error {
	return tr.db.Create(token).Error
}

func (tr *TokenRepository) Update(token *models.Token) error {
	return tr.db.Save(token).Error
}

func (tr *TokenRepository) Delete(token *models.Token) error {
	return tr.db.Delete(token).Error
}

func (tr *TokenRepository) FindAll() ([]models.Token, error) {
	var tokens []models.Token
	if err := tr.db.Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}
