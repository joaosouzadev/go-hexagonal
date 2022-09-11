package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joaosouzadev/go-hexagonal-arch/internal/adapters/db"
	"github.com/joaosouzadev/go-hexagonal-arch/internal/application"
	"github.com/joaosouzadev/go-hexagonal-arch/pkg/app"
	"github.com/joaosouzadev/go-hexagonal-arch/pkg/utils"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productOutputDto application.ProductOutputDto
	productOutputDto.HydrateFromEntity(product)

	c.JSON(http.StatusOK, productOutputDto)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	uuid := c.Param("uuid")
	if valid := utils.IsValidUUID(uuid); valid != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	product, err := h.productService.Get(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productOutputDto application.ProductOutputDto
	productOutputDto.HydrateFromEntity(product)

	c.JSON(http.StatusOK, productOutputDto)
}
