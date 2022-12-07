package usuario

import "tradingcards/infrastructure/persistence/usuario"

type IUsuario interface {
	BuscarUsuarioPorEmail(email *string) (res *usuario.User, err error)
	AdicionarUsuario(req *usuario.User) (id *int64, err error)
	AlterarUsuario(req *usuario.User) (err error)
	DeletarUsuario(id *int64) (err error)
}
