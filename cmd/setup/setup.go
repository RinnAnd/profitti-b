package setup

import (
	"database/sql"
	"os"
	"profitti/internal/app/transport/http/handlers/categories"
	"profitti/internal/app/transport/http/handlers/expenses"
	e_handler "profitti/internal/app/transport/http/handlers/expenses"
	f_handler "profitti/internal/app/transport/http/handlers/financials"
	"profitti/internal/app/transport/http/handlers/partnership"
	handler "profitti/internal/app/transport/http/handlers/partnership"
	"profitti/internal/app/transport/http/handlers/users"
	uc "profitti/internal/core/usecases/categories"
	e_usecase "profitti/internal/core/usecases/expenses"
	"profitti/internal/core/usecases/financials"
	"profitti/internal/core/usecases/login"
	p_usecase "profitti/internal/core/usecases/partnership"
	cat_repo "profitti/internal/infra/database/repository/categories"
	expense "profitti/internal/infra/database/repository/expenses"
	"profitti/internal/infra/database/repository/financial"
	p_repo "profitti/internal/infra/database/repository/partnership"
	"profitti/internal/infra/database/repository/user"
	"profitti/internal/infra/service/auth"
	srv_cat "profitti/internal/infra/service/categories"
	e_service "profitti/internal/infra/service/expenses"
	f_service "profitti/internal/infra/service/financials"
	p_service "profitti/internal/infra/service/partnership"
	service "profitti/internal/infra/service/users"
	"time"
)

type Setup struct {
	RegisterHandler            users.RegisterHandler
	LoginHandler               users.LoginHandler
	CreateFinancialHandler     f_handler.CreateHandler
	GetFinancialsByUserHandler f_handler.GetByUserHandler
	CreateExpenseHandler       expenses.CreateHandler
	GetExpensesByUserHandler   expenses.GetByUserHandler
	CreatePartnership          partnership.CreateHandler
	GetPartnerships            partnership.GetHandler
	CreateCategory             categories.Handler
	GetCategories              categories.GetHandler
}

func Build(db *sql.DB) *Setup {

	secret := os.Getenv("JWT_SECRET")

	userRepository := user.New(db)
	userService := service.New(userRepository)
	userHandler := users.NewRegister(userService)

	partnershipRepo := p_repo.New(db)
	partnershipSrv := p_service.New(partnershipRepo)

	partnershipUsecase := p_usecase.New(partnershipSrv, userService)
	partnershipHandler := handler.NewCreateHandler(partnershipUsecase)
	getPartnershipHandler := handler.NewGet(partnershipUsecase)

	auth := auth.New(secret, time.Minute*40)
	loginusecase := login.New(userService, auth)

	loginHandler := users.NewLogin(loginusecase)

	financialRepository := financial.New(db)
	financialService := f_service.New(financialRepository)

	createUseCase := financials.NewCreateUseCase(financialService)
	getByUserUseCase := financials.NewGetUserFinancials(financialService)

	createHandler := f_handler.NewCreate(createUseCase)
	getByUserHandler := f_handler.NewGetByUser(getByUserUseCase)

	expensesRepository := expense.New(db)
	expensesService := e_service.New(expensesRepository, financialRepository, partnershipRepo)

	e_createUseCase := e_usecase.NewCreateUseCase(expensesService)
	e_getByUserUseCase := e_usecase.NewGetByUserUseCase(expensesService, userService)

	e_createHandler := e_handler.NewCreate(e_createUseCase)
	e_getByUserHandler := e_handler.NewGetByUser(e_getByUserUseCase)

	c_repository := cat_repo.New(db)
	catService := srv_cat.New(c_repository)

	createcatusecase := uc.New(userService, catService)
	getcatusecase := uc.NewG(userService, catService)
	createcathandler := categories.New(createcatusecase)

	getcategorieshandler := categories.NewG(getcatusecase)

	setup := &Setup{
		RegisterHandler:            userHandler,
		LoginHandler:               loginHandler,
		CreatePartnership:          partnershipHandler,
		GetPartnerships:            getPartnershipHandler,
		CreateFinancialHandler:     createHandler,
		GetFinancialsByUserHandler: getByUserHandler,
		CreateExpenseHandler:       e_createHandler,
		GetExpensesByUserHandler:   e_getByUserHandler,
		CreateCategory:             createcathandler,
		GetCategories:              getcategorieshandler,
	}

	return setup
}
