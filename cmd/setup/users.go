package setup

import (
	"database/sql"
	"profitti/internal/app/transport/http/handlers/users"
	"profitti/internal/infra/database/repository/user"
	service "profitti/internal/infra/service/users"
)

func User(db *sql.DB, stp *Setup) {
	userRepository := user.New(db)
	userService := service.New(userRepository)
	userHandler := users.NewRegister(userService)
	loginHandler := users.NewLogin(userService)

	stp.RegisterHandler = userHandler
	stp.LoginHandler = loginHandler
}
