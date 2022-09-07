package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joaosouzadev/go-hexagonal-arch/adapters/db"
	"github.com/joaosouzadev/go-hexagonal-arch/application"
	"github.com/joaosouzadev/go-hexagonal-arch/pkg/app"
	"net/http"
)

type ProductHandler struct {
	app            *app.App
	productService application.ProductServiceInterface
}

func NewProductHandler(app *app.App) *ProductHandler {
	return &ProductHandler{
		app:            app,
		productService: application.NewProductService(db.NewProductDb(app.DBConn)),
	}
}

func (h *ProductHandler) NewProduct(c *gin.Context) {
	var productInputDto application.ProductInputDto
	if err := c.ShouldBindJSON(&productInputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.productService.Create(productInputDto)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productOutputDto application.ProductOutputDto
	productOutputDto.HydrateFromEntity(product)

	c.JSON(http.StatusOK, productOutputDto)
}
