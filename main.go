package main

import (
	"github.com/joaosouzadev/go-hexagonal-arch/internal/adapters/server"
	"github.com/joaosouzadev/go-hexagonal-arch/pkg/app"
)

func main() {
	application := app.NewApp()
	api := server.NewAPIServer(application)
	api.Start()
}
