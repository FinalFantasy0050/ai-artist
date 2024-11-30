package usermanager

import (
	"ai-artist/gateway/utils/logging/logDefault"
	"encoding/json"
	"os"
)

func Init() {
	loadUsers()
}

func loadUsers() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		logDefault.Error("There is no file. (" + fileName + ")")
		os.Exit(1)
	}

	var userStruct users
	if err := json.Unmarshal(data, &userStruct); err != nil {
		logDefault.Error("JSON unmarshal error.")
		os.Exit(1)
	}

	usermap = make(map[string]bool)
	for _, user := range userStruct.User {
		usermap[user] = true
	}
}
