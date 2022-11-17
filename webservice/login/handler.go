package login

import (
	"net/http"
	"tradingcards/application/login"

	"github.com/gin-gonic/gin"
)

// faz login e retorna token
func logar(c *gin.Context) {
	var req login.Login

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := login.Logar(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"token": token})
}
