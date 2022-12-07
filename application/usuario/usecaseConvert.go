package usuario

import usuarioData "tradingcards/infrastructure/persistence/usuario"

func ConveterUsuarioToData(req *UserReq) (res *usuarioData.User) {
	res = &usuarioData.User{
		ID:    req.ID,
		Nome:  req.Nome,
		Email: req.Email,
		Senha: req.Senha,
		CPF:   req.CPF,
	}
	return
}
