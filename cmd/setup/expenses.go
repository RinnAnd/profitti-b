package setup

import (
	"database/sql"
	handler "profitti/internal/app/transport/http/handlers/expenses"
	"profitti/internal/core/usecases/expenses"
	expense "profitti/internal/infra/database/repository/expenses"
	service "profitti/internal/infra/service/expenses"
)

func Expenses(db *sql.DB, stp *Setup) {
	expensesRepository := expense.New(db)
	expensesService := service.New(expensesRepository)

	createUseCase := expenses.NewCreateUseCase(expensesService)
	getByUserUseCase := expenses.NewGetByUserUseCase(expensesService)

	createHandler := handler.NewCreate(createUseCase)
	getByUserHandler := handler.NewGetByUser(getByUserUseCase)

	stp.CreateExpenseHandler = createHandler
	stp.GetExpensesByUserHandler = getByUserHandler
}
