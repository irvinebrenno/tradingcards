package main

import (
	"tradingcards/webservice/login"
	"tradingcards/webservice/usuario"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// definimos os grupos de rotas
	usuarioGroup := r.Group("")
	loginGroup := r.Group("")

	usuario.Router(usuarioGroup)
	login.Router(loginGroup)

	r.Run() // rodando em localhost:8080
}
