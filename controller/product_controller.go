package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	if idParam == "" {
		response := model.Response{
			Message: "id is required",
			Code:    http.StatusBadRequest,
		}
		ctx.JSON(response.Code, gin.H{
			"error": response,
		})
		return
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		response := model.Response{
			Message: "id must be a number",
			Code:    http.StatusBadRequest,
		}

		ctx.JSON(response.Code, gin.H{
			"error": response,
		})
		return
	}

	product, err := p.productUsecase.GetProductByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if product == nil {
		response := model.Response{
			Message: "product not found",
			Code:    http.StatusNotFound,
		}

		ctx.JSON(response.Code, gin.H{
			"error": response,
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}
