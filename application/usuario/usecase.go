package usuario

import (
	"fmt"
	"tradingcards/config/database"
	"tradingcards/domain/usuario"
	"tradingcards/services"
)

func AdicionarUsuario(req *UserReq) (id *int64, err error) {
	mensagemPadrao := "Erro ao cadastrar usuário: "

	db := database.Conectar()
	defer db.Close()

	repoUsuario := usuario.NewRepo(db)

	senha256 := services.SHA256Encoder(*req.Senha)

	req.Senha = &senha256

	dados := ConveterUsuarioToData(req)

	id, err = repoUsuario.AdicionarUsuario(dados)
	if err != nil {
		return nil, fmt.Errorf("%v %v", mensagemPadrao, err.Error())
	}

	return
}

func AlterarUsuario(req *UserReq) (err error) {
	mensagemPadrao := "Erro ao alterar usuário: "

	db := database.Conectar()
	defer db.Close()

	repoUsuario := usuario.NewRepo(db)

	senha256 := services.SHA256Encoder(*req.Senha)

	req.Senha = &senha256

	dados := ConveterUsuarioToData(req)

	err = repoUsuario.AlterarUsuario(dados)
	if err != nil {
		return fmt.Errorf("%v %v", mensagemPadrao, err.Error())
	}

	return
}

func DeletarUsuario(id *int64) (err error) {
	mensagemPadrao := "Erro ao deletar usuário: "

	db := database.Conectar()
	defer db.Close()

	repoUsuario := usuario.NewRepo(db)

	err = repoUsuario.DeletarUsuario(id)
	if err != nil {
		return fmt.Errorf("%v %v", mensagemPadrao, err.Error())
	}

	return
}
