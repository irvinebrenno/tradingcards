package usuario

import "tradingcards/infrastructure/persistence/usuario"

type IUsuario interface {
	BuscarUsuarioPorEmail(email *string) (res *usuario.User, err error)
}
