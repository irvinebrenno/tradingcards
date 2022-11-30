package usuario

type User struct {
	ID    *int64  `banco:"id"`
	Nome  *string `banco:"nome"`
	Senha *string `banco:"senha"`
	Email *string `banco:"email"`
	CPF   *int64  `banco:"cpf"`
}
