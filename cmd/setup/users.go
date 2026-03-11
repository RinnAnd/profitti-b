package setup

import (
	"database/sql"
	"os"
	"profitti/internal/app/transport/http/handlers/users"
	"profitti/internal/core/usecases/login"
	"profitti/internal/infra/database/repository/user"
	"profitti/internal/infra/service/auth"
	service "profitti/internal/infra/service/users"
	"time"
)

func User(db *sql.DB, stp *Setup) {
	secret := os.Getenv("JWT_SECRET")

	userRepository := user.New(db)
	userService := service.New(userRepository)
	userHandler := users.NewRegister(userService)

	auth := auth.New(secret, time.Minute*10)
	loginusecase := login.New(userService, auth)

	loginHandler := users.NewLogin(loginusecase)

	stp.RegisterHandler = userHandler
	stp.LoginHandler = loginHandler
}
