package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/flugika/seller-web-app/entity"
	"github.com/gin-gonic/gin"
)

// POST /product
func CreateProduct(c *gin.Context) {
	var product entity.Product
	var productType entity.ProductType
	var promotion entity.Promotion
	var user entity.User

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา productType ด้วย id
	if tx := entity.DB().Where("id = ?", product.ProductTypeID).First(&productType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "productType not found"})
		return
	}

	// ค้นหา promotion ด้วย id
	if tx := entity.DB().Where("id = ?", product.PromotionID).First(&promotion); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "promotion not found"})
		return
	}

	// ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", product.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// สร้าง Product
	cp := entity.Product{
		Name:         product.Name,
		ProductPhoto: product.ProductPhoto,
		Description:  product.Description,
		ProductType:  productType,
		Promotion:    promotion,
		User:         user,
	}

	// บันทึก
	if err := entity.DB().Create(&cp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": cp})
}

// GET /product/:id
func GetProduct(c *gin.Context) {
	var product entity.Product
	id := c.Param("id")

	if tx := entity.DB().Preload("ProductType").Preload("Promotion").Preload("User").Where("id = ?", id).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /products
func ListProducts(c *gin.Context) {
	var products []entity.Product
	if err := entity.DB().Preload("ProductType").Preload("Promotion").Preload("User").Raw("SELECT * FROM products").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// DELETE /product/:id
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM products WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /product
func UpdateProduct(c *gin.Context) {
	var product entity.Product
	var productType entity.ProductType
	var promotion entity.Promotion
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา productType ด้วย id
	if tx := entity.DB().Where("id = ?", product.ProductTypeID).First(&productType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "productType not found"})
		return
	}

	// ค้นหา promotion ด้วย id
	if tx := entity.DB().Where("id = ?", product.PromotionID).First(&promotion); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "promotion not found"})
		return
	}

	// ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", product.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// อัพเดท Product
	up := entity.Product{
		Name:         product.Name,
		ProductPhoto: product.ProductPhoto,
		Description:  product.Description,
		ProductType:  productType,
		Promotion:    promotion,
		User:         user,
	}

	// บันทึก
	if err := entity.DB().Where("id = ?", product.ID).Updates(&up).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": up})

}
