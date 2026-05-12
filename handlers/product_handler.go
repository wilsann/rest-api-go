package handlers

import (
	"net/http"
	"rest-api-go/domain/request"
	"rest-api-go/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	products, err := models.GetList()
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

func ProductDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	product, err := models.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get product. %v" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"product": product,
	})
}

func ProductCreate(c *gin.Context) {
	var req request.ProductCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	layout := "2006-01-02 15:04:05"
	timeStr := time.Now().Format(layout)
	t, _ := time.Parse(layout, timeStr)

	userId := c.GetInt64("userId")
	p := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
		CreatedAt:   t,
		CreatedBy:   userId,
	}

	err := p.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed create product.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"product": p,
	})
}

func ProductUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	userId := c.GetInt64("userId")
	product, err := models.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get product. %v" + err.Error(),
		})
		return
	}

	if product.CreatedBy != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var updatedProduct models.Product
	err = c.ShouldBindJSON(&updatedProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body.",
		})
		return
	}

	updatedProduct.ID = id
	err = updatedProduct.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed update product.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"product": updatedProduct,
	})
}

func ProductDelete(c *gin.Context) {
	userId := c.GetInt64("userId")
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	product, err := models.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get product. %v" + err.Error(),
		})
		return
	}

	if product.CreatedBy != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	err = product.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed delete product.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
