package controllers

import (
	"go-api/models"
	"go-api/usecases"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecases.ProductUsecase
}

func NewProductController(productUsecase usecases.ProductUsecase) productController {
	return productController{productUsecase}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()

	if err != nil {
		log.Fatalf("Error get products... %v", err)
		ctx.JSON(500, gin.H{
			"error": gin.H{
				"message": "Error get products",
			},
		})
	}

	ctx.JSON(200, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}

	product, err := p.productUsecase.GetProductById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
