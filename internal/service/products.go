package service 

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"product_service/internal/db"
	"fmt"
)

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing product ID")
		return
	}

	var product, err = db.GetProduct(id)
	if err != nil {
		respondError(c, 404, "Product not found")
		return
	}
	c.JSON(200, product)
}

func GetProducts(c *gin.Context) {
	offset, err1 := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, err2 := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if err1 != nil || err2 != nil {
		respondError(c, 400, "Offset and limit must be integers")
		return
	}
	if limit < 0 || offset < 0 {
		respondError(c, 400, "Offset and limit must be non-negative")
		return
	}

	limit = min(limit, 200)
	offset *= limit

	products, err := db.GetProducts(limit, offset)
	if err != nil {
		fmt.Println(err)
		respondError(c, 404, "Products not found")
		return
	}

	if len(products) == 0 {
		c.JSON(200, []any{})
		return
	}
	c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
	var product db.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		respondError(c, 400, "Invalid JSON")
		return
	}

	id, err := db.CreateProduct(&product)
	if err != nil {
		respondError(c, 400, "Failed to create product")
		return
	}

	product.Id = &id
	c.JSON(201, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing product ID")
		return
	}

	err := db.DeleteProduct(id)
	if err != nil {
		respondError(c, 404, "Product not found or already deleted")
		return
	}

	respondSuccess(c, 200, "Product deleted")
}
