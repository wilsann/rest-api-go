package handlers

import (
	"net/http"
	"rest-api-go/models"
	"rest-api-go/utils"
	"strconv"

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
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err = product.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed create product.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"product": product,
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

	_, err = models.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get product. %v" + err.Error(),
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
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	result, err := models.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get product. %v" + err.Error(),
		})
		return
	}

	err = result.Delete()
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
