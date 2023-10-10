package handler

import (
	"Reusable-Auth-System/pkg/users"
	"context"
	"fmt"
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
	corsConfig.AllowOrigins = []string{"*", "http://localhost:8080"}

	// Registering MiddleWares
	handler.Router.Use(cors.New(corsConfig))
	handler.Router.Use(gin.Logger())
	handler.Router.Use(gin.Recovery())

	handler.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome 🚀",
		})
	})

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
		v1.GET("/username/:username", h.GetByUsername)
		v1.GET("/:username", h.GetByUsername)
		v1.GET("/full_name/:full_name", h.GetUserByFullName)
		v1.PUT("/:id", h.UpdateUserByID)
		v1.PUT("/:id", h.SetActivity)
	}

}

func AliveCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Alive!",
	})
}

// Serve - gracefully serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		err := h.Router.Run(fmt.Sprintf(":%v", 8080))
		if err != nil {
			log.Println(err)

		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	// CreateAccount a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := h.Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down gracefully")
	return nil
}
