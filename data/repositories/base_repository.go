package repositories

import (
	"github.com/google/uuid"
)

type ResourceType string

const (
	UserType ResourceType = "user"
)

type Repository interface {
	FindByIDAndUserID(resourceID uuid.UUID, userID uuid.UUID) (interface{}, error)
}

var RepositoryMap = map[ResourceType]Repository{
	UserType: NewUserRepository(),
}
