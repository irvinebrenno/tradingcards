package usuario

import (
	"tradingcards/config/database"
	"tradingcards/domain/usuario"
)

func AdicionarUsuario(req *UserReq) (err error) {
	mensagemPadrao := "Erro ao cadastrar usuário: "

	db := database.Conectar()
	defer db.Close()

	repoUsuario := usuario.NewRepo(db)

	return
}
