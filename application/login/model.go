package login

type Login struct {
	ID    *int64
	Email *string `json:"email"`
	Senha *string `json:"senha"`
}
