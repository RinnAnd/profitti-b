package setup

import (
	"database/sql"
	"profitti/internal/app/transport/http/handlers/users"
	"profitti/internal/infra/database/repository/user"
	service "profitti/internal/infra/service/users"
)

func User(db *sql.DB) *Setup {
	userRepository := user.New(db)
	userService := service.New(userRepository)
	userHandler := users.NewUserHandler(userService)

	return &Setup{
		UserHandler: userHandler,
	}
}
