package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type VoteRepository struct {
	db *gorm.DB
}

func NewVoteRepository() *VoteRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &VoteRepository{db: db}
}

func (vr *VoteRepository) FindByID(id uuid.UUID) (*models.Vote, error) {
	var m models.Vote
	if evr := vr.db.First(&m, id).Error; evr != nil {
		return nil, evr
	}
	return &m, nil
}

func (vr *VoteRepository) FindAll() ([]models.Vote, error) {
	var votes []models.Vote
	if evr := vr.db.Find(&votes).Error; evr != nil {
		return nil, evr
	}
	return votes, nil
}

func (vr *VoteRepository) Create(m *models.Vote) error {
	return vr.db.Create(m).Error
}

func (vr *VoteRepository) Update(m *models.Vote) error {
	return vr.db.Save(m).Error
}

func (vr *VoteRepository) Delete(m *models.Vote) error {
	return vr.db.Delete(m).Error
}

func (vr *VoteRepository) DeleteByID(id uuid.UUID) bool {
	evr := vr.db.Delete(&models.Vote{}, id)
	if evr != nil {
		return true
	}
	return false
}
