package models

import "errors"

type User struct {
	ID           int    `json:"id,omitempty"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

// NOTE: This is just for demo purposes, in real scenario the user would be searched for in the DB
var users = []User{
	{
		ID:           1,
		Username:     "testuser",
		PasswordHash: "testpass", // Not really a hash, demo purposes only
	},
}

func AuthenticateUser(username string, password string) (*User, error) {
	for _, user := range users {
		if user.Username == username && user.PasswordHash == password {
			return &user, nil
		}
	}

	err := errors.New("authentication failed")
	return nil, err
}

func GetUserById(id int) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
