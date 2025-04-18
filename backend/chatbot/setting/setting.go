package setting

import (
	"ai-artist/chatbot/utils/logging/logDefault"
	"encoding/json"
	"os"
)

func Init() {
	err := readSettingFile()
	if err != nil {
		logDefault.Error(err)

		os.Exit(1)
	}
	logDefault.System("Successfully finished initializing setting.")
}

func readSettingFile() error {
	file, err := os.ReadFile(settingFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &Setting)
	if err != nil {
		return err
	}

	return nil
}
