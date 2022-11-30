package postgres

import (
	"database/sql"
	"fmt"
	"tradingcards/infrastructure/persistence/usuario"
)

type BDUsuario struct {
	DB *sql.DB
}

// BuscarUsuarioPorEmail busca um usuário pelo email no banco de dados
func (postgres *BDUsuario) BuscarUsuarioPorEmail(email *string) (res *usuario.User, err error) {
	res = &usuario.User{}

	if err := postgres.DB.QueryRow(
		`
		SELECT
			TU.id,
			TU.nome,
			TU.email,
			TU.cpf,
			TU.senha
		FROM t_usuario TU
		WHERE
			TS.sequence ilike $1
		`, email).Scan(&res.ID, &res.Nome, &res.Email, &res.CPF, &res.Senha); err != nil {
		return nil, fmt.Errorf("Erro ao buscar usuário: %v", err.Error())
	}

	return
}
