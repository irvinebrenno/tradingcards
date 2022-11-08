package postgres

import (
	"database/sql"
)

type BDUsuario struct {
	DB *sql.DB
}

// InsertSequence insere um usuario no banco de dados
func (postgres *BDUsuario) InsertSequence() (err error) {
	return nil
}
