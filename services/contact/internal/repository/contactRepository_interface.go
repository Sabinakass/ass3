package repository

import "architecture_go/services/contact/internal/domain"

type ContactRepository interface {
	Create(contact *domain.Contact) error

	Retrieve(id int) (*domain.Contact, error)

	Update(contact *domain.Contact) error

	Delete(id int) error

	CreateGroup(group *domain.Group) error

	RetrieveGroup(id int) (*domain.Group, error)

	AddContactToGroup(contactID, groupID int) error
}
