package controllers

import (
	"deathtiny_encounters/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GoogleLogin(c *gin.Context) {
	url := services.GetGoogleOAuthURL()
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleAuthCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := services.HandleGoogleCallback(code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
