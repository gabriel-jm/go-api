package main

import (
	"go-api/controllers"
	"go-api/db"
	"go-api/repositories"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	server := gin.Default()

	productRepository := repositories.NewProductRepository(dbConnection)
	productUsecase := usecases.NewProductUsecase(productRepository)
	productController := controllers.NewProductController(productUsecase)
	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	server.GET("/products/:id", productController.GetProductById)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "Pong")
	})

	server.Run(":8000")
}
