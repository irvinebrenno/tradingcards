package usuario

import (
	"net/http"
	"tradingcards/application/usuario"

	"github.com/gin-gonic/gin"
)

func AdicionarUsuario(c *gin.Context) {
	var req usuario.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := usuario.AdicionarUsuario(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"dados": id})
}
