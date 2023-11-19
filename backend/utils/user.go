package utils

import "backend/database"

func CreatePreferences(preferences map[string]int, tags []string) map[string]int {
	newPreferences := make(map[string]int, len(tags))
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

func AddPreferenceValue(user *database.User, tags []string, value int) {
	for _, tag := range tags {
		user.Preferences[tag] += value
	}
}
