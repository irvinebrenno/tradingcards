package usuario

import (
	"database/sql"
	"tradingcards/infrastructure/persistence/usuario"
	"tradingcards/infrastructure/persistence/usuario/postgres"
)

// repository -
type repository struct {
	Data *postgres.BDUsuario
}

// newRepo -
func newRepo(newDB *sql.DB) *repository {
	return &repository{
		Data: &postgres.BDUsuario{DB: newDB},
	}
}

// BuscarUsuarioPorEmail é um gerenciador de fluxo para buscar um usuário pelo email no banco de dados
func (r *repository) BuscarUsuarioPorEmail(email *string) (res *usuario.User, err error) {
	return r.Data.BuscarUsuarioPorEmail(email)
}
