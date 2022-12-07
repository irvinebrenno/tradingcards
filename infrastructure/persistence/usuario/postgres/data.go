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

// AdicionarUsuario adiciona um usuário no banco de dados
func (postgres *BDUsuario) AdicionarUsuario(req *usuario.User) (id *int64, err error) {
	id = new(int64)

	if err := postgres.DB.QueryRow(
		`
		INSERT INTO t_usuario (nome, email, cpf, senha)
    	VALUES ($1::TEXT, $2::TEXT, $3::BIGINT, $4::TEXT)
		RETURNING id
		`, req.Nome, req.Email, req.CPF, req.Senha).Scan(&id); err != nil {
		return nil, fmt.Errorf("Erro ao adicionar usuário: %v", err.Error())
	}

	return
}

// AlterarUsuario altera um usuário no banco de dados
func (postgres *BDUsuario) AlterarUsuario(req *usuario.User) (err error) {

	sqlStatement := `UPDATE t_usuario
	SET nome=$2::TEXT, email=$3::TEXT, cpf=$4::BIGINT, senha=$5::TEXT
	WHERE id=$1::BIGINT`
	_, err = postgres.DB.Exec(sqlStatement, req.ID, req.Nome, req.Email, req.CPF, req.Senha)
	if err != nil {
		return err
	}

	return
}

// DeletarUsuario um usuário no banco de dados
func (postgres *BDUsuario) DeletarUsuario(id *int64) (err error) {

	sqlStatement := `DELETE FROM t_usuario
	WHERE id=$1::BIGINT`
	_, err = postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return
}

func (postgres *BDUsuario) ListarUsuarios() (res []usuario.User, err error) {

	var user usuario.User

	rows, err := postgres.DB.
		Query(`
			SELECT
				tu.id,
				tu.nome,
				tu.email,
				tu.cpf
			FROM t_usuario tu`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Nome, &user.Email, &user.CPF)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}

	return
}
