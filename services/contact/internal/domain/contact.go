package domain

import (
	"errors"
	"unicode"
)

type Contact struct {
	ID          int
	firstName   string
	lastName    string
	middleName  string
	PhoneNumber string
	FullName    string
}

func NewContact(id int, firstName, lastName, middleName, phoneNumber string) (*Contact, error) {
	if !isOnlyDigits(phoneNumber) {
		return nil, errors.New("phone number must contain only digits")
	}

	return &Contact{
		ID:          id,
		firstName:   firstName,
		lastName:    lastName,
		middleName:  middleName,
		PhoneNumber: phoneNumber,
		FullName:    lastName + " " + firstName + " " + middleName,
	}, nil
}

func (c *Contact) GetFullName() string {
	return c.lastName + " " + c.firstName + " " + c.middleName
}

func isOnlyDigits(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
