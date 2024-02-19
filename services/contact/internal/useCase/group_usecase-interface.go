package usecase

import "architecture_go/services/contact/internal/domain"

type GroupUseCase interface {
	CreateGroup(group *domain.Group) error

	GetGroup(id int) (*domain.Group, error)

	AddContactToGroup(contactID, groupID int) error
}
