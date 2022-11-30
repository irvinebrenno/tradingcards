package usuario

// UserReq define uma estrutura para adicionar usuários
type UserReq struct {
	ID    *int64  `json:"id"`
	Nome  *string `json:"nome"`
	Senha *string `json:"senha"`
	Email *string `json:"email"`
	CPF   *int64  `json:"cpf"`
}
