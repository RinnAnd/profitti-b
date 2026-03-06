package setup

import (
	"database/sql"
	"profitti/internal/app/transport/http/handlers/users"
)

type Setup struct {
	UserHandler users.UserHandler
}

func Build(db *sql.DB) *Setup {
	setup := &Setup{}
	User(db, setup)

	return setup
}
