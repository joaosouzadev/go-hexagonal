package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joaosouzadev/go-hexagonal-arch/adapters/server/handler"
	"github.com/joaosouzadev/go-hexagonal-arch/pkg/app"
	"log"
)

type APIServer struct {
	httpServer     *gin.Engine
	productHandler *handler.ProductHandler
}

func NewAPIServer(app *app.App) *APIServer {
	return &APIServer{
		httpServer:     gin.Default(),
		productHandler: handler.NewProductHandler(app),
	}
}

func (a *APIServer) Start() {
	a.httpServer.POST("/products", a.productHandler.NewProduct)

	err := a.httpServer.Run(":9000")
	if err != nil {
		log.Fatal(err)
	}
}
