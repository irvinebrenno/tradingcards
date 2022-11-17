package login

import (
	"fmt"
	"tradingcards/services"
)

func Logar(req *Login) (token *string, err error) {

	str := "senhaSHA256"

	// aqui vamos comparar as senhas criptografadas após implementar o banco
	if str != services.SHA256Encoder(*req.Senha) {
		return nil, fmt.Errorf("Senha incorreta")
	}

	tokenStr, err := services.NewJWTService().GenerateToken(uint(*req.ID))
	if err != nil {
		return nil, fmt.Errorf("Credencial inválida")
	}

	token = &tokenStr

	return
}
