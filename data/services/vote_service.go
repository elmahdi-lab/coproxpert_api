package services

import (
	"errors"

	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/repositories"
)

func CreateVote(r *models.Vote) (*models.Vote, error) {
	voteRepository := repositories.NewVoteRepository()
	if voteRepository == nil {
		return nil, errors.New("error creating vote repository")
	}

	err := voteRepository.Create(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetVoteByID(id uuid.UUID) (*models.Vote, error) {
	voteRepository := repositories.NewVoteRepository()
	if voteRepository == nil {
		return nil, errors.New("error creating vote repository")
	}

	vote, err := voteRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return vote, nil
}

func UpdateVote(r *models.Vote) (*models.Vote, error) {
	voteRepository := repositories.NewVoteRepository()
	if voteRepository == nil {
		return nil, errors.New("error creating vote repository")
	}

	err := voteRepository.Update(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func DeleteVoteByID(id uuid.UUID) bool {
	voteRepository := repositories.NewVoteRepository()
	if voteRepository == nil {
		return false
	}

	return voteRepository.DeleteByID(id)
}
