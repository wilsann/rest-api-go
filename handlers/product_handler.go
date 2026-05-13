package handlers

import (
	"net/http"
	"rest-api-go/domain/request"
	"rest-api-go/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	ProductList(c *gin.Context)
	ProductDetail(c *gin.Context)
	ProductCreate(c *gin.Context)
	ProductUpdate(c *gin.Context)
	ProductDelete(c *gin.Context)
}

type ProductHandlerInteractor struct {
	usecase usecase.ProductUsecase
}

func ProductHandlerImpl(productUsecase usecase.ProductUsecase) ProductHandler {
	return &ProductHandlerInteractor{
		usecase: productUsecase,
	}
}

func (hi *ProductHandlerInteractor) ProductList(c *gin.Context) {
	products, err := hi.usecase.ProductList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get products. %v" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"products": products,
	})
}

func (hi *ProductHandlerInteractor) ProductDetail(c *gin.Context) {
	productId, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	product, err := hi.usecase.ProductDetail(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"product": product,
	})
}

func (hi *ProductHandlerInteractor) ProductCreate(c *gin.Context) {
	var req request.ProductCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	userId := c.GetInt64("userId")

	data, err := hi.usecase.ProductCreate(req, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed create product.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"product": data,
	})
}

func (hi *ProductHandlerInteractor) ProductUpdate(c *gin.Context) {
	var req request.ProductUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	userId := c.GetInt64("userId")

	data, err := hi.usecase.ProductUpdate(req, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"product": data,
	})
}

func (hi *ProductHandlerInteractor) ProductDelete(c *gin.Context) {
	userId := c.GetInt64("userId")
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	err = hi.usecase.ProductDelete(id, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get product. %v" + err.Error(),
		})
		return
	}

}
