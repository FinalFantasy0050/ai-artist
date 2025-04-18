package handler

import (
	"net/http"

	"github.com/unrolled/render"
)

type Handler struct {
	http.Handler
}
