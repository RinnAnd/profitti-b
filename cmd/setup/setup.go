package setup

import (
	"database/sql"
	"profitti/internal/app/transport/http/handlers/expenses"
	"profitti/internal/app/transport/http/handlers/financials"
	"profitti/internal/app/transport/http/handlers/partnership"
	"profitti/internal/app/transport/http/handlers/users"
)

type Setup struct {
	// Users
	RegisterHandler users.RegisterHandler
	LoginHandler    users.LoginHandler
	// Financials
	CreateFinancialHandler     financials.CreateHandler
	GetFinancialsByUserHandler financials.GetByUserHandler
	// Expenses
	CreateExpenseHandler     expenses.CreateHandler
	GetExpensesByUserHandler expenses.GetByUserHandler
	// Partnerships
	CreatePartnership partnership.CreateHandler
	GetPartnerships   partnership.GetHandler
}

func Build(db *sql.DB) *Setup {
	setup := &Setup{}
	User(db, setup)
	Financials(db, setup)
	Expenses(db, setup)
	Partnership(db, setup)

	return setup
}
