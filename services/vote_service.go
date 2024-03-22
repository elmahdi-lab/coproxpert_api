package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateVote(r *models.Vote) (*models.Vote, error) {
	voteRepository, err := repositories.NewVoteRepository()
	if err != nil {
		return nil, err
	}

	err = voteRepository.Create(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetVoteByID(id uuid.UUID) (*models.Vote, error) {
	voteRepository, err := repositories.NewVoteRepository()
	if err != nil {
		return nil, err
	}

	vote, err := voteRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return vote, nil
}

func UpdateVote(r *models.Vote) (*models.Vote, error) {
	voteRepository, err := repositories.NewVoteRepository()
	if err != nil {
		return nil, err
	}

	err = voteRepository.Update(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func DeleteVoteByID(id uuid.UUID) bool {
	voteRepository, err := repositories.NewVoteRepository()
	if err != nil {
		return false
	}

	return voteRepository.DeleteByID(id)
}
