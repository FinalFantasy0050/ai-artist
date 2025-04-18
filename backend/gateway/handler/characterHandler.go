package handler

import (
	"ai-artist/gateway/setting"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

const inferCharacterMethod = "/infer/character"

type imageResponse struct {
	Image string `json:"image"`
}

type chatCharacterResponse struct {
	Text string `json:"text"`
}

type characterResponse struct {
	Image string `json:"image"`
	Text  string `json:"text"`
}

func (h *Handler) characterHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	defer r.Body.Close()

	if strings.TrimSpace(req.Prompt) == "" {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	exist := usermanager.Authentication(req.User)
	if !exist {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	err = logToFile(req.User, req.Prompt)
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	url := setting.Setting.ChatbotServer + inferCharacterMethod
	respChat, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	url = setting.Setting.ImageGeneratorServer + inferImageMethod
	respImg, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		rend.JSON(w, http.StatusBadGateway, nil)
		return
	}

	var imageResp imageResponse
	err = json.NewDecoder(respImg.Body).Decode(&imageResp)
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	defer respImg.Body.Close()

	var chatCharacterResp chatCharacterResponse
	err = json.NewDecoder(respChat.Body).Decode(&chatCharacterResp)
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	defer respChat.Body.Close()

	resp := characterResponse{
		Image: imageResp.Image,
		Text:  chatCharacterResp.Text,
	}

	rend.JSON(w, http.StatusOK, resp)
}
