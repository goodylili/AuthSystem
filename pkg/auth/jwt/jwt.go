package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
)

func GenerateAccessJWT(username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(4 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = username
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AuthMiddleWare is a function that returns a Gin middleware handler.
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token string from the "access_token" cookie
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			// If there's an error (e.g., the cookie is not present), abort the request and respond with an error
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			return
		}

		// jwt.Parse is used to parse the JWT token string obtained from the cookie.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check if the signing method used in the token matches the expected signing method (HS256 in this case).
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, jwt.ErrSignatureInvalid // Return an error if the signing method does not match.
			}
			// Return the key used for signing the token. This key is obtained from an environment variable.
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// If there's an error in parsing the token or if the token is invalid, abort the request and respond with an error.
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Assert that the claims in the token are of type jwt.MapClaims, which is a map of strings to interfaces.
		claims, ok := token.Claims.(jwt.MapClaims)
		// If the assertion fails or the token is invalid, abort the request and respond with an error.
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
			return
		}

		// Obtain the expiration time claim from the token and assert it to a float64 (since it's a Unix timestamp).
		expiry, ok := claims["exp"].(float64)
		// If the assertion fails or the expiration time is in the past, abort the request and respond with an error.
		if !ok || time.Unix(int64(expiry), 0).Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		// Set the username claim from the token in the Gin context for use in subsequent handlers.
		c.Set("username", claims["user"])
		// Call c.Next() to pass the request to the next handler in the chain.
		c.Next()
	}
}
