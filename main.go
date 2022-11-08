package main

import (
	"tradingcards/webservice/usuario"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// definimos os grupos de rotas
	usuarioGroup := r.Group("")

	usuario.Router(usuarioGroup)

	r.Run() // rodando em localhost:8080
}
