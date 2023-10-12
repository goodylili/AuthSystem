package social

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type GoogleUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Prevents CSRF attacks during GoogleOAuthInstance Authentication
var (
	oauthStateString    = uuid.New().String()
	GoogleOAuthInstance = oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     getEnVars()[0],
		ClientSecret: getEnVars()[1],
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
)

// HandleGoogleLogin function is a handler for the route where users initiate the GoogleOAuthInstance login process.
// It redirects users to the GoogleOAuthInstance login page from my page.
func HandleGoogleLogin(c *gin.Context) {
	url := GoogleOAuthInstance.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleGoogleCallback function is a handler for the route where users are redirected after they have authenticated with GoogleOAuthInstance.
func HandleGoogleCallback(c *gin.Context) {
	content, err := GetGoogleUserInfo(c.Query("state"), c.Query("code"))
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.String(http.StatusOK, "Content: %s\n", content)
}

// GetGoogleUserInfo function gets the user info from GoogleOAuthInstance
func GetGoogleUserInfo(state string, code string) (GoogleUser, error) {
	var user GoogleUser

	if state != oauthStateString {
		return user, fmt.Errorf("invalid oauth state")
	}

	token, err := GoogleOAuthInstance.Exchange(context.TODO(), code)
	if err != nil {
		return user, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return user, fmt.Errorf("failed creating request: %s", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	response, err := httpClient.Do(req)
	if err != nil {
		return user, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return user, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	err = json.Unmarshal(contents, &user)
	if err != nil {
		return user, fmt.Errorf("failed unmarshalling user info: %s", err.Error())
	}

	return user, nil
}
