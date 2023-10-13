package handler

import (
	"Reusable-Auth-System/pkg/auth/jwt"
	"Reusable-Auth-System/pkg/auth/social"
	"Reusable-Auth-System/pkg/users"
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router *gin.Engine
	Users  users.Service
	Server *http.Server
}

// Initialize - gracefully serves our newly set up handler function
func (h *Handler) Initialize() error {
	// Initialize the Gin server in a goroutine
	go func() {
		if err := h.Server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("ListenAndServe: %v\n", err)
		}
	}()

	// Listen for a termination signal (SIGTERM)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := h.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

	log.Println("Server shut down gracefully")
	return nil
}

// NewHandler - returns a pointer to a Handler
func NewHandler(users users.Service) *Handler {
	log.Println("setting up our handlers")
	handler := &Handler{
		Users: users,
	}

	// Create a new Gin router
	handler.Router = gin.Default()

	corsConfig := cors.DefaultConfig()

	// corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Authorization", "Content-Length", "Content-Type"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}

	// Registering MiddleWares
	handler.Router.Use(cors.New(corsConfig))
	handler.Router.Use(gin.Logger())
	handler.Router.Use(gin.Recovery())

	handler.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome 🚀",
		})
	})

	// Initialize your HTTP server with the Gin router
	handler.Server = &http.Server{
		Addr: ":8080", // Set the server address
	}

	return handler
}

func (h *Handler) mapRoutes() {
	h.Router.GET("/alive", AliveCheck)
	h.Router.GET("/ready", h.Ping)

	// Users Routes
	v1 := h.Router.Group("/api/v1/users")
	{
		v1.POST("/", h.CreateUser)
		v1.GET("/:id", h.GetUserByID)
		v1.GET("/email/:email", h.GetByEmail)
		v1.GET("/username/:username", jwt.AuthMiddleWare(), h.GetByUsername)
		v1.GET("/:username", h.GetByUsername)
		v1.GET("/full_name/:full_name", jwt.AuthMiddleWare(), h.GetUserByFullName)
		v1.PUT("/:id", jwt.AuthMiddleWare(), h.UpdateUserByID)
		v1.PUT("/:id", h.SetActivity)

		v1.POST("/sign_in", h.SignIn)
		v1.POST("/sign_out", jwt.AuthMiddleWare(), h.SignOut)
		v1.GET("/google/sign_in", social.HandleGoogleLogin)
		v1.GET("/google/callback", social.HandleGoogleCallback)
		v1.GET("/github/sign_in", social.HandleGitHubLogin)
		v1.GET("/github/callback", social.HandleGitHubCallback)

	}

}

func AliveCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Alive!",
	})
}
