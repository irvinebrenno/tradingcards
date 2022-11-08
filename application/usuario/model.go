package usuario

type UserReq struct {
	Nome  *string `json:"nome"`
	Senha *string `json:"senha"`
	Email *string `json:"email"`
}
