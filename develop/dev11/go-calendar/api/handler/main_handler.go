package handler

import "net/http"

// GET /
func (h *Handler) MainHandler(_ *http.Request) APIResponse {
	return h.JSON(http.StatusOK, map[string]string{"main": "/ server route"})
}
