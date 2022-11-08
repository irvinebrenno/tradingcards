package sequence

type UserReq struct {
	Nome  *string `banco:"nome"`
	Senha *string `banco:"senha"`
	Email *string `banco:"email"`
}
