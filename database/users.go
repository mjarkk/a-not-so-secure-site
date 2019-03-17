package database

import (
	"fmt"
	"strconv"
)

// UserT is the data type for a user
type UserT struct {
	Password string
	Username string
	ID       string
}

// SeedUsers ads some data to the database
func SeedUsers() error {
	_, err := NewUser("admin", "a-cool-password")
	if err != nil {
		return err
	}
	_, err = NewUser("lol", "123")
	return err
}

// Users returns a list of all users
func Users() ([]UserT, error) {
	toReturn := []UserT{}
	rows, err := DB.Query("SELECT username, password, ID FROM users")
	if err != nil {
		return toReturn, err
	}

	for rows.Next() {
		toAdd := UserT{}
		err := rows.Scan(&toAdd.Username, &toAdd.Password, &toAdd.ID)
		if err != nil {
			return toReturn, err
		}
		toReturn = append(toReturn, toAdd)
	}

	return toReturn, nil
}

// User returns a user
func User(whereWhat, is string) (UserT, error) {
	_, err := strconv.ParseInt(is, 10, 64)
	if err != nil {
		is = "\"" + is + "\""
	}

	toReturn := UserT{}
	row := DB.QueryRow("SELECT username, password, ID FROM users WHERE " + whereWhat + " = " + is)
	err = row.Scan(&toReturn.Username, &toReturn.Password, &toReturn.ID)

	return toReturn, err
}

// NewUser creates a new user
func NewUser(username, password string) (UserT, error) {
	query, err := DB.Prepare("INSERT INTO users (`username`, `password`) VALUES (\"" + username + "\", \"" + password + "\")")
	if err != nil {
		return UserT{}, err
	}
	res, err := query.Exec()
	if err != nil {
		return UserT{}, err
	}

	id, err := res.LastInsertId()
	return UserT{
		Username: username,
		Password: password,
		ID:       fmt.Sprintf("%v", id),
	}, err
}
