package handler

import (
	"authenticationsystem/pkg/database"
	"authenticationsystem/pkg/users"
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

func (handler *Handler) mapRoutes() {
	handler.Router.GET("/alive", AliveCheck)
	handler.Router.GET("/ready", ReadyCheck)

	// Users Routes
	v1 := handler.Router.Group("/api/v1")
	{
		v1.POST("/user", CreateUser)
		v1.GET("/users/:id", GetUserByID)
		v1.GET("/users/email/:email", GetByEmail)
		v1.GET("/users/username/:username", GetByUsername)
		v1.GET("/users/:username", GetByUsername)
		v1.GET("/users/fullname/:fullname", GetUserByFullName)
		v1.PUT("/users/:id", UpdateUserByID)
		v1.PUT("/users/:id", DeactivateUserByID)
	}

}

func AliveCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Alive!",
	})
}

// Serve - gracefully serves our newly set up handler function
func (handler *Handler) Serve() error {
	go func() {
		err := handler.Router.Run(fmt.Sprintf(":%v", 8080))
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
	err := handler.Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down gracefully")
	return nil
}

func ReadyCheck(c *gin.Context) {
	err := database.ReadyCheck(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Database is ready"})
}
