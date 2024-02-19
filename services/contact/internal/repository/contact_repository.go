package repository

import (
	"architecture_go/services/contact/internal/domain"
	"database/sql"
	"errors"
	"strings"
)

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(contact *domain.Contact) error {
	parts := strings.Fields(contact.GetFullName())
	FirstName := ""
	LastName := ""
	MiddleName := ""
	switch len(parts) {
	case 0:
		return errors.New("full name cannot be empty")
	case 1:
		FirstName = parts[0]
	case 2:
		FirstName = parts[0]
		LastName = parts[1]
	default:
		FirstName = parts[0]
		LastName = parts[len(parts)-1]
		MiddleName = strings.Join(parts[1:len(parts)-1], " ")
	}

	_, err := r.db.Exec("INSERT INTO contacts (first_name, last_name, middle_name, phone_number) VALUES ($1, $2, $3, $4)",
		FirstName, LastName, MiddleName, contact.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *contactRepository) Retrieve(id int) (*domain.Contact, error) {
	row := r.db.QueryRow("SELECT id, full_name, phone_number FROM contacts WHERE id = $1", id)
	contact := domain.Contact{}
	err := row.Scan(&contact.ID, &contact.FullName, &contact.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactRepository) Update(contact *domain.Contact) error {
	parts := strings.Fields(contact.GetFullName())
	FirstName := ""
	LastName := ""
	MiddleName := ""
	switch len(parts) {
	case 0:
		return errors.New("full name cannot be empty")
	case 1:
		FirstName = parts[0]
	case 2:
		FirstName = parts[0]
		LastName = parts[1]
	default:
		FirstName = parts[0]
		LastName = parts[len(parts)-1]
		MiddleName = strings.Join(parts[1:len(parts)-1], " ")
	}
	_, err := r.db.Exec("UPDATE contacts SET first_name = $1, last_name = $2, middle_name = $3, phone_number = $4 WHERE id = $5",
		FirstName, LastName, MiddleName, contact.PhoneNumber, contact.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *contactRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM contacts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *contactRepository) CreateGroup(group *domain.Group) error {
	_, err := r.db.Exec("INSERT INTO groups (name) VALUES ($1)", group.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *contactRepository) RetrieveGroup(id int) (*domain.Group, error) {
	row := r.db.QueryRow("SELECT id, name FROM groups WHERE id = $1", id)
	group := domain.Group{}
	err := row.Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *contactRepository) AddContactToGroup(contactID, groupID int) error {
	_, err := r.db.Exec("INSERT INTO contacts_groups (contact_id, group_id) VALUES ($1, $2)", contactID, groupID)
	if err != nil {
		return err
	}
	return nil
}
