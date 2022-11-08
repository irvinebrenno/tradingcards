package usuario

import (
	"database/sql"
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
