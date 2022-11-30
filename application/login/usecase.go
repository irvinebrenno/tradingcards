package login

import (
	"fmt"
	"tradingcards/config/database"
	"tradingcards/domain/usuario"
	"tradingcards/services"
)

func Logar(req *Login) (token *string, err error) {
	mensagemPadrao := "Erro ao logar: "

	db := database.Conectar()
	defer db.Close()

	repoUsuario := usuario.NewRepo(db)

	user, err := repoUsuario.BuscarUsuarioPorEmail(req.Email)
	if err != nil {
		return nil, err
	}

	// aqui vamos comparar as senhas criptografadas após implementar o banco
	if *user.Senha != services.SHA256Encoder(*req.Senha) {
		return nil, fmt.Errorf("%v Senha incorreta", mensagemPadrao)
	}

	tokenStr, err := services.NewJWTService().GenerateToken(uint(*req.ID))
	if err != nil {
		return nil, fmt.Errorf("%v Credencial inválida", mensagemPadrao)
	}

	token = &tokenStr

	return
}
