package controllers

import (
	"net/http"

	"github.com/mrojasb2000/fullstack/api/responses"
)

// Home - home handler
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
