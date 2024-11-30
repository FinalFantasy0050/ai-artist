package handler

import (
	"ai-artist/gateway/setting"
	usermanager "ai-artist/gateway/userManager"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const inferWriterMethod = "/infer/writer"

func (h *Handler) aiWriterHandler(w http.ResponseWriter, r *http.Request) {
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

	url := setting.Setting.ChatbotServer + inferWriterMethod
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	defer resp.Body.Close()

	rend.Data(w, http.StatusOK, respBody)
}
