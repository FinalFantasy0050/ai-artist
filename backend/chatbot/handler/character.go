package handler

import (
	openapi "ai-artist/chatbot/openAPI"
	"ai-artist/chatbot/utils/logging"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) characterHandler(w http.ResponseWriter, r *http.Request) {
	var request requestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	prompt := characterPromptEngineering(request.Prompt)
	logging.Logger.Info(request.User + " : " + prompt)

	outputPrompt, inputToken, outputToken, totalToken := openapi.Chatbot(prompt)
	logToFile(request.User, "Character", request.Prompt, outputPrompt, inputToken, outputToken, totalToken)

	response := responseBody{
		Text: outputPrompt,
	}

	rend.JSON(w, http.StatusOK, response)
}

func characterPromptEngineering(prompt string) string {
	const role = "너는 베스트셀러 소설 작가야. 창의적이고 매력적인 캐릭터의 스토리를 작성해줘."
	const purpose = "다음 입력을 바탕으로 캐릭터의 설명을 작성해."
	const promptStruct = "캐릭터의 설명은 다음 입력을 바탕으로 묘사해줘. 깔끔하게 설명만 출력해."
	prePrompt := fmt.Sprintf("%s\n%s\n%s\n", role, purpose, promptStruct)

	return fmt.Sprintf("%s```%s```", prePrompt, prompt)
}
