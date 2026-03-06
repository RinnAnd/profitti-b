package main

import (
	"os"
	"profitti/internal/app/transport/server"
)

func main() {
	server := server.StartServer(os.Getenv("PORT"))
	server.Run()
}
