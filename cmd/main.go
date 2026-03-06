package main

import (
	"os"
	"profitti/cmd/setup"
	"profitti/internal/app/transport/http/routes"
	"profitti/internal/app/transport/server"
	"profitti/internal/infra/database/connection"
)

func main() {
	cfg := connection.New(os.Getenv("GOOSE_DBSTRING"))
	db := cfg.Open()
	setup := setup.Build(db)

	rtr := routes.Routes{
		UserHandler: setup.UserHandler,
	}

	server := server.StartServer(os.Getenv("PORT"))
	server.Run()
}
