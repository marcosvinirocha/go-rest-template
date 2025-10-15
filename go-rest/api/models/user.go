package models

import (
	"errors"
	"sync"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Mock database
var (
	users = []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	}
	nextID = 3
	mutex  = &sync.Mutex{}
)

// GetAllUsers returns all users
func GetAllUsers() []User {
	mutex.Lock()
	defer mutex.Unlock()
	return users
}

// GetUserByID returns a user by ID
func GetUserByID(id int) (User, error) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

// CreateUser creates a new user
func CreateUser(user User) User {
	mutex.Lock()
	defer mutex.Unlock()
	user.ID = nextID
	nextID++
	users = append(users, user)
	return user
}

// UpdateUser updates an existing user
func UpdateUser(id int, updatedUser User) (User, error) {
	mutex.Lock()
	defer mutex.Unlock()
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			return updatedUser, nil
		}
	}
	return User{}, errors.New("user not found")
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	mutex.Lock()
	defer mutex.Unlock()
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}