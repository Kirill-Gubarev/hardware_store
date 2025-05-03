package service 

import (
	"github.com/gin-gonic/gin"
	"user_service/internal/db"
	"time"
)

type VerifyRequest struct {
	Login string `json:"login"`
	Token string `json:"token"`
}
func VerifySession(c *gin.Context){
	var verifyRequest VerifyRequest
	err := c.ShouldBindJSON(&verifyRequest)
	if err != nil {
		respondError(c, 400, "Invalid JSON")
		return
	}
	session, err := db.GetSession(verifyRequest.Login, verifyRequest.Token)
	if err != nil {
		c.JSON(401, gin.H{"role":nil})
		return
	}

	if time.Since(*session.CreatedAt) > 24 * time.Hour {
		db.DeleteSession(*session.Id)
		c.JSON(401, gin.H{"role":nil})
		return
	}

	db.UpdateSession(*session.Id)
	c.JSON(200, gin.H{"role" : session.User.Role.Name})
}

type LoginRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}
func LoginSession(c *gin.Context) {
	var loginRequest LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		respondError(c, 400, "Invalid JSON")
		return
	}

	token, err := db.CreateSession(loginRequest.Login, loginRequest.Password)
	if err != nil {
		respondError(c, 400, "Failed to create session")
		return
	}

	c.JSON(201, gin.H{"token":token})
}

func LogoutSession(c *gin.Context) {
	var verifyRequest VerifyRequest
	err := c.ShouldBindJSON(&verifyRequest)
	if err != nil {
		respondError(c, 400, "Invalid JSON")
		return
	}
	session, err := db.GetSession(verifyRequest.Login, verifyRequest.Token)
	if err != nil {
		c.JSON(401, gin.H{"access" : false})
		return
	}

	err = db.DeleteSession(*session.Id)
	if err != nil {
		respondError(c, 404, "Session not found or already deleted")
		return
	}

	respondSuccess(c, 200, "Session deleted")
}
