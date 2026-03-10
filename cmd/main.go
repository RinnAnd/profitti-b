package main

import (
	"os"
	"profitti/cmd/setup"
	"profitti/internal/app/transport/http/routes"
	"profitti/internal/app/transport/server"
	"profitti/internal/infra/database/connection"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		print(err.Error())
		return
	}
	cfg := connection.New(os.Getenv("GOOSE_DBSTRING"))
	db := cfg.Open()
	setup := setup.Build(db)

	rtr := routes.Routes{
		RegisterHandler: setup.RegisterHandler,
		LoginHandler:    setup.LoginHandler,
	}

	server := server.StartServer(os.Getenv("PORT"))
	rtr.Init(server.G)
	server.Run()
}
