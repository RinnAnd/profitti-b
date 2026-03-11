package setup

import (
	"database/sql"
	"profitti/internal/app/transport/http/handlers/financials"
	"profitti/internal/app/transport/http/handlers/users"
)

type Setup struct {
	// Users
	RegisterHandler users.RegisterHandler
	LoginHandler    users.LoginHandler
	// Financials
	CreateFinancialHandler financials.CreateHandler
	GetByUserHandler       financials.GetByUserHandler
}

func Build(db *sql.DB) *Setup {
	setup := &Setup{}
	User(db, setup)
	Financials(db, setup)

	return setup
}
