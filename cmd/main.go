package main

import (
	"os"
	"profitti/cmd/setup"
	"profitti/internal/app/transport/http/routes"
	"profitti/internal/app/transport/server"
	"profitti/internal/infra/database/connection"
	"profitti/internal/infra/service/auth"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		print(err.Error())
		return
	}
	cfg := connection.New(os.Getenv("GOOSE_DBSTRING"))
	secret := os.Getenv("JWT_SECRET")

	db := cfg.Open()
	setup := setup.Build(db)

	rtr := routes.Routes{
		RegisterHandler:            setup.RegisterHandler,
		LoginHandler:               setup.LoginHandler,
		CreateFinancialHandler:     setup.CreateFinancialHandler,
		GetFinancialsByUserHandler: setup.GetFinancialsByUserHandler,
		CreateExpenseHandler:       setup.CreateExpenseHandler,
		GetExpensesByUserHandler:   setup.GetExpensesByUserHandler,
		CreatePartnership:          setup.CreatePartnership,
		GetPartnerships:            setup.GetPartnerships,
		CreateCategory:             setup.CreateCategory,
		GetCategories:              setup.GetCategories,
	}

	auth := auth.New(secret, time.Minute*40)

	server := server.StartServer(os.Getenv("PORT"), auth)
	rtr.Init(server.G)
	server.Run()
}
