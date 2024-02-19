package repository

import (
	"architecture_go/services/contact/internal/domain"
	"database/sql"
)

type groupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) GroupRepository {
	return &groupRepository{
		db: db,
	}
}

func (r *groupRepository) CreateGroup(group *domain.Group) error {
	_, err := r.db.Exec("INSERT INTO groups (name) VALUES ($1)", group.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *groupRepository) RetrieveGroup(id int) (*domain.Group, error) {
	row := r.db.QueryRow("SELECT id, name FROM groups WHERE id = $1", id)
	group := &domain.Group{}
	err := row.Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (r *groupRepository) AddContactToGroup(contactID, groupID int) error {
	_, err := r.db.Exec("INSERT INTO contacts_groups (contact_id, group_id) VALUES ($1, $2)", contactID, groupID)
	if err != nil {
		return err
	}
	return nil
}
