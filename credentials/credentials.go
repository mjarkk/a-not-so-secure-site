package credentials

import (
	"github.com/mjarkk/a-not-so-secure-site/database"
	"github.com/mjarkk/a-not-so-secure-site/utils"
)

var sessions = map[string]database.UserT{}

// CreateSession creates a new sessions
func CreateSession(user database.UserT) (string, error) {
	key := ""
	for {
		rand, err := utils.RandomString(20)
		if err != nil {
			return "", err
		}
		if _, ok := sessions[rand]; !ok {
			key = rand
			break
		}
	}
	sessions[key] = user
	return key, nil
}

// GetSession returns a userobject from a session if provided the right key
func GetSession(key string) (database.UserT, bool) {
	user, ok := sessions[key]
	return user, ok
}
