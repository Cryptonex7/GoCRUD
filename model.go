package main

import (
	"database/sql"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

func (u *User) getUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *User) updateUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *User) deleteUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (u *User) createUser(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func getUsers(db *sql.DB, start, count int) ([]User, error) {
	return nil, errors.New("Not Implemented")
}
