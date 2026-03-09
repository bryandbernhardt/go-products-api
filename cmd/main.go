package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	// camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/product", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.GetProductById)

	server.POST("/product", ProductController.CreateProduct)

	server.Run(":8000")
}
