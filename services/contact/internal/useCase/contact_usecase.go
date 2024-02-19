package usecase

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
	"errors"
	"fmt"
)

type contactUseCase struct {
	contactRepo repository.ContactRepository
	groupRepo   repository.GroupRepository
}

func NewContactUseCase(contactRepo repository.ContactRepository, groupRepo repository.GroupRepository) ContactUseCase {
	return &contactUseCase{
		contactRepo: contactRepo,
		groupRepo:   groupRepo,
	}
}

func (c *contactUseCase) CreateContact(contact *domain.Contact) error {
	if contact == nil {
		return errors.New("contact cannot be nil")
	}
	if contact.FullName == "" || contact.PhoneNumber == "" {
		return errors.New("name and phone number must be provided")
	}
	return c.contactRepo.Create(contact)
}

func (c *contactUseCase) GetContact(id int) (*domain.Contact, error) {
	if id <= 0 {
		return nil, errors.New("invalid contact ID")
	}
	return c.contactRepo.Retrieve(id)
}

func (c *contactUseCase) UpdateContact(contact *domain.Contact) error {
	if contact == nil {
		return errors.New("contact cannot be nil")
	}
	if contact.ID <= 0 || contact.FullName == "" || contact.PhoneNumber == "" {
		return errors.New("valid ID, name, and phone number must be provided")
	}
	_, err := c.GetContact(contact.ID)
	if err != nil {
		return fmt.Errorf("contact does not exist: %w", err)
	}
	return c.contactRepo.Update(contact)
}

func (c *contactUseCase) DeleteContact(id int) error {
	if id <= 0 {
		return errors.New("invalid contact ID")
	}
	return c.contactRepo.Delete(id)
}

func (c *contactUseCase) CreateGroup(group *domain.Group) error {
	if group == nil {
		return errors.New("group cannot be nil")
	}
	if len(group.Name) == 0 || len(group.Name) > 250 {
		return errors.New("group name must be between 1 and 250 characters")
	}
	return c.groupRepo.CreateGroup(group)
}

func (c *contactUseCase) GetGroup(id int) (*domain.Group, error) {
	if id <= 0 {
		return nil, errors.New("invalid group ID")
	}
	return c.groupRepo.RetrieveGroup(id)
}

func (c *contactUseCase) AddContactToGroup(contactID, groupID int) error {
	if contactID <= 0 || groupID <= 0 {
		return errors.New("both contact and group IDs must be valid")
	}
	_, err := c.GetContact(contactID)
	if err != nil {
		return fmt.Errorf("contact does not exist: %w", err)
	}
	_, err = c.GetGroup(groupID)
	if err != nil {
		return fmt.Errorf("group does not exist: %w", err)
	}
	return c.groupRepo.AddContactToGroup(contactID, groupID)
}
