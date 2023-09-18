package handler

import (
	"authenticationsystem/pkg/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router *gin.Engine
	Users  user.UserService
	Server *http.Server
}

// NewHandler - returns a pointer to a Handler
func NewHandler(users user.UserService) *Handler {
	log.Println("setting up our handlers")
	h := &Handler{
		Users: users,
	}

	// Create a new Gin router
	h.Router = gin.Default()

	h.Server = &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router, // Use Gin router as the handler
	}

	return h
}
