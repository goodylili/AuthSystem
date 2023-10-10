package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	// Assume the user is authenticated
	session := sessions.Default(c)
	session.Set("authenticated", true)
	err := session.Save()
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"message": "authenticated"})
}

func secureEndpoint(c *gin.Context) {
	if !isLoggedIn(c) {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}
	c.JSON(200, gin.H{"message": "Hello, authorized user!"})
}

func isLoggedIn(c *gin.Context) bool {
	session := sessions.Default(c)
	authenticated := session.Get("authenticated")
	return authenticated != nil && authenticated.(bool)
}
