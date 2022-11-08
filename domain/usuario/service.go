package usuario

import "database/sql"

// NewRepo retorna um novo repositório com um contexo de banco
func NewRepo(DB *sql.DB) IUsuario {
	return newRepo(DB)
}
