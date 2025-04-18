package handler

import (
	openapi "ai-artist/chatbot/openAPI"
	"ai-artist/chatbot/utils/logging"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type requestBody struct {
	Prompt string `json:"prompt"`
	User   string `json:"user"`
}

type responseBody struct {
	Text string `json:"text"`
}

func (h *Handler) writerHandler(w http.ResponseWriter, r *http.Request) {
	var request requestBody
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	defer r.Body.Close()

	prompt := writerPromptEngineering(request.Prompt)
	logging.Logger.Info(request.User + " : " + prompt)

	outputPrompt, inputToken, outputToken, totalToken := openapi.Chatbot(prompt)
	logToFile(request.User, "AI Writer", request.Prompt, outputPrompt, inputToken, outputToken, totalToken)

	response := responseBody{
		Text: outputPrompt,
	}

	rend.JSON(w, http.StatusOK, response)
}

func writerPromptEngineering(prompt string) string {
	const role = "너는 베스트셀러 소설 작가야. 창의적이고 매력적인 이야기를 작성해줘."
	const purpose = "다음 입력을 바탕으로 단편 소설을 작성해."
	const promptStruct = "스토리는 도입부, 전개, 절정, 결말의 구조로 매끄럽게 작성해. 스토리 글 이외에는 작성하지마. 도입부, 전개, 절정, 결말을 언급하지마."
	prePrompt := fmt.Sprintf("%s\n%s\n%s\n", role, purpose, promptStruct)

	return fmt.Sprintf("%s```%s```", prePrompt, prompt)
}

func logToFile(user string, aiType string, inputPrompt string, outputPrompt string, inputToken int, outputToken int, totalToken int) error {
	logDir := "./log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("fale to create log directory: %v", err)
		}
	}

	logFilePath := filepath.Join(logDir, fmt.Sprintf("%s.txt", user))
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	logContent := fmt.Sprintf(
		"Time: %s\nType: %s\nInput: %s\nOutput: %s\nPrompt Tokens: %d\nCompletion Tokens: %d\nTotal Tokens: %d\n\n",
		currentTime, aiType, inputPrompt, outputPrompt, inputToken, outputToken, totalToken,
	)

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(logContent); err != nil {
		return fmt.Errorf("failed to write to log file: %v", err)
	}

	return nil
}
