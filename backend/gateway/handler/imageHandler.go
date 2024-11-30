package handler

import (
	"ai-artist/gateway/setting"
	usermanager "ai-artist/gateway/userManager"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const inferImageMethod = "/infer/image"

func (h *Handler) imageHandler(w http.ResponseWriter, r *http.Request) {
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

	/*
		forwardReq, err := http.NewRequest("POST", setting.Setting.ImageGeneratorServer+inferImageMethod, bytes.NewBuffer(jsonData))
		if err != nil {
			rend.JSON(w, http.StatusBadRequest, nil)
			return
		}
		forwardReq.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(forwardReq)
		if err != nil {
			rend.JSON(w, http.StatusBadRequest, "client do error.")
			return
		}
		defer resp.Body.Close()
	*/
	url := setting.Setting.ImageGeneratorServer + inferImageMethod
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

func logToFile(user string, prompt string) error {
	const logDir = "./log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create log directory: %v", err)
		}
	}

	logFilePath := filepath.Join(logDir, fmt.Sprintf("%s.txt", user))
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	logContent := fmt.Sprintf(
		"Time: %s\nInput Prompt: %s\n\n",
		currentTime,
		prompt,
	)

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(logContent)
	if err != nil {
		return fmt.Errorf("failed to write to log file: %v", err)
	}

	return nil
}
