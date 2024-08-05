package services

import (
	"context"
	"deathtiny_encounters/config"
	"deathtiny_encounters/models"
	"deathtiny_encounters/repositories"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	google_oauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"time"
)

var (
	oauthConfig = &oauth2.Config{
		RedirectURL:  config.GetOAuthConfig()["redirectURL"],
		ClientID:     config.GetOAuthConfig()["clientID"],
		ClientSecret: config.GetOAuthConfig()["clientSecret"],
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
)

func GetGoogleOAuthURL() string {
	return oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func HandleGoogleCallback(code string) (string, error) {
	ctx := context.Background()
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	oauth2Service, err := google_oauth2.NewService(ctx, option.WithTokenSource(oauthConfig.TokenSource(ctx, token)))
	if err != nil {
		return "", err
	}

	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		return "", err
	}

	user, err := repositories.GetUserByEmail(userInfo.Email)
	if err != nil {
		// If user not found, create new user
		user = models.User{
			Email:    userInfo.Email,
			Username: userInfo.Name,
		}
		if err := repositories.CreateUser(&user); err != nil {
			return "", err
		}
	}

	// Generate JWT token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
