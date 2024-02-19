package domain

import "errors"

type Group struct {
	ID   int
	Name string
}

func NewGroup(id int, name string) (*Group, error) {
	if len(name) > 250 {
		return nil, errors.New("group name exceeds the maximum length of 250 characters")
	}
	return &Group{
		ID:   id,
		Name: name,
	}, nil
}
