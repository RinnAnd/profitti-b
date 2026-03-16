package setup

import (
	"database/sql"
	handler "profitti/internal/app/transport/http/handlers/financials"
	"profitti/internal/core/usecases/financials"
	"profitti/internal/infra/database/repository/financial"
	service "profitti/internal/infra/service/financials"
)

func Financials(db *sql.DB, stp *Setup) {
	financialRepository := financial.New(db)
	financialService := service.New(financialRepository)

	createUseCase := financials.NewCreateUseCase(financialService)
	getByUserUseCase := financials.NewGetUserFinancials(financialService)

	createHandler := handler.NewCreate(createUseCase)
	getByUserHandler := handler.NewGetByUser(getByUserUseCase)

	stp.CreateFinancialHandler = createHandler
	stp.GetFinancialsByUserHandler = getByUserHandler
}
