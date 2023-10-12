package social

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io/ioutil"
	"net/http"
)

var (
	GitHubOAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     getEnVars()[2],
		ClientSecret: getEnVars()[3],
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
)

type GitHubUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

func HandleGitHubLogin(c *gin.Context) {
	url := GitHubOAuthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGitHubCallback(c *gin.Context) {
	user, err := GetUserInfo(c.Query("state"), c.Query("code"))
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserInfo(state string, code string) (GitHubUser, error) {
	var user GitHubUser

	if state != oauthStateString {
		return user, fmt.Errorf("invalid oauth state")
	}

	token, err := GitHubOAuthConfig.Exchange(context.TODO(), code)
	if err != nil {
		return user, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return user, fmt.Errorf("failed creating request: %s", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return user, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	err = json.Unmarshal(contents, &user)
	if err != nil {
		return user, fmt.Errorf("failed unmarshalling user info: %s", err.Error())
	}

	return user, nil
}
