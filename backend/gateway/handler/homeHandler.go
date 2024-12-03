package handler

import "net/http"

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "login.html", http.StatusTemporaryRedirect)
}
