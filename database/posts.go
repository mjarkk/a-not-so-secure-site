package database

import (
	"fmt"
	"strconv"
)

// PostT is the data type for a post
type PostT struct {
	Title   string
	Content string
	User    UserT
	ID      string
}

// SeedPosts ads some data to the database
func SeedPosts() error {
	err := NewPost("Welcom", "This is a unsecure site :), sql injections and xss work everywhere here", 2)
	if err != nil {
		return err
	}
	err = NewPost("users and there passwords", "admin -> a-cool-password<br>lol -> 123<br>Not as if they are fully visable in the database", 1)
	return err
}

// Posts returns a list of all users
func Posts() ([]PostT, error) {
	toReturn := []PostT{}
	rows, err := DB.Query("SELECT title, content, userID, ID FROM posts")
	if err != nil {
		return toReturn, err
	}

	for rows.Next() {
		toAdd := PostT{}
		var userID uint32
		err := rows.Scan(&toAdd.Title, &toAdd.Content, &userID, &toAdd.ID)
		if err != nil {
			return toReturn, err
		}

		user, err := User("ID", fmt.Sprintf("%v", userID))
		if err != nil {
			return toReturn, err
		}
		toAdd.User = user

		toReturn = append(toReturn, toAdd)
	}

	return toReturn, nil
}

// Post returns a user
func Post(whereWhat, is string) (PostT, error) {
	_, err := strconv.ParseInt(is, 10, 64)
	if err != nil {
		is = "`" + is + "`"
	}

	toReturn := PostT{}
	row := DB.QueryRow("SELECT title, content, userID, ID FROM posts WHERE " + whereWhat + " = " + is)

	var userID uint32
	err = row.Scan(&toReturn.Title, &toReturn.Content, &userID, &toReturn.ID)

	user, err := User("ID", fmt.Sprintf("%v", userID))
	if err != nil {
		return toReturn, err
	}
	toReturn.User = user

	return toReturn, err
}

// NewPost creates a new user
func NewPost(title, content string, userID uint32) error {
	query := fmt.Sprintf(
		"INSERT INTO posts (title, content, userID) VALUES (\"%v\", \"%v\", %v)",
		title,
		content,
		userID,
	)
	_, err := DB.Query(query)
	return err
}
