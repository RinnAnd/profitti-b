package setup

import (
	"database/sql"
	"profitti/internal/app/transport/http/handlers/users"
)

type Setup struct {
	RegisterHandler users.RegisterHandler
	LoginHandler    users.LoginHandler
}

func Build(db *sql.DB) *Setup {
	setup := &Setup{}
	User(db, setup)

	return setup
}
