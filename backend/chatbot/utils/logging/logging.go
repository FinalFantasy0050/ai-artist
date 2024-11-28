package logging

import "ai-artist/chatbot/utils/logging/loggingIPFS"

func Init() {
	Logger = newLogger()

	Logger.Init()
}

func newLogger() Logging {
	return loggingIPFS.NewLogger()
}
