package repository

import "architecture_go/services/contact/internal/domain"

type GroupRepository interface {
	CreateGroup(group *domain.Group) error
	RetrieveGroup(id int) (*domain.Group, error)
	AddContactToGroup(contactID, groupID int) error
}
