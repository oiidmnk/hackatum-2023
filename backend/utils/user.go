package utils

import "backend/database"

func CreatePreferences(preferences map[string]uint32, tags []string) map[string]uint32 {
	newPreferences := make(map[string]uint32, len(tags))
	for _, tag := range tags {
		if val, ok := preferences[tag]; !ok {
			newPreferences[tag] = 100
		} else {
			newPreferences[tag] = val
		}
	}
	return newPreferences
}

func SetUserPreferences(users []*database.User, tags []string) {
	for _, user := range users {
		user.Preferences = CreatePreferences(user.Preferences, tags)
	}
}
