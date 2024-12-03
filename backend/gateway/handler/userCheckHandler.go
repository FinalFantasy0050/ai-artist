package handler

import (
	usermanager "ai-artist/gateway/userManager"
	"encoding/json"
	"net/http"
	"strings"
)

func (h *Handler) userCheckHandler(w http.ResponseWriter, r *http.Request) {
	var req userRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	defer r.Body.Close()

	if strings.TrimSpace(req.User) == "" {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	exist := usermanager.Authentication(req.User)
	if !exist {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	rend.JSON(w, http.StatusOK, nil)
}
