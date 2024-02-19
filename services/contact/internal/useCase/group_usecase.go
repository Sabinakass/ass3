package usecase

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
	"errors"
)

type groupUseCase struct {
	groupRepo repository.GroupRepository
}

func NewGroupUseCase(gr repository.GroupRepository) GroupUseCase {
	return &groupUseCase{
		groupRepo: gr,
	}
}

func (g *groupUseCase) CreateGroup(group *domain.Group) error {
	if group.Name == "" {
		return errors.New("the group name cannot be empty")
	}
	return g.groupRepo.CreateGroup(group)
}

func (g *groupUseCase) GetGroup(id int) (*domain.Group, error) {
	if id <= 0 {
		return nil, errors.New("invalid group ID")
	}
	return g.groupRepo.RetrieveGroup(id)
}

func (g *groupUseCase) AddContactToGroup(contactID, groupID int) error {
	if contactID <= 0 || groupID <= 0 {
		return errors.New("contactID and groupID must be valid non-zero integers")
	}
	return g.groupRepo.AddContactToGroup(contactID, groupID)
}
