// models.user.go
// Initializes application's user model

package main

import (
	"errors"
	"strings"
)

type user struct {
	Username string `json:username`
	Password string `json:password`
	Name     string `json:name`
}

var userList = []user{
	user{Username: "user1", Password: "pass1", Name: "Sean"},
	user{Username: "user2", Password: "pass2", Name: "Sam"},
	user{Username: "user3", Password: "pass3", Name: "John"},
}

// Check for validity of user login input
// Later can add password strength, etc.
func isUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func registerNewUser(username, password, name string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	} else if strings.TrimSpace(name) == "" {
		return nil, errors.New("the name can't be empty")
	}

	u := user{Username: username, Password: password, Name: name}

	userList = append(userList, u)

	return &u, nil
}

func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}

	return true
}

func getAccountName(username, password string) string {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return u.Name
		}
	}

	return ""
}
