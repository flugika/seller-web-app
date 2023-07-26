package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/flugika/seller-web-app/entity"
	"github.com/gin-gonic/gin"
)

// POST /product
func CreatePurchase(c *gin.Context) {
	var purchase entity.Purchase
	var paymentMethod entity.PaymentMethod
	var receipt entity.Receipt
	var user entity.User
	var product entity.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา paymentMethod ด้วย id
	if tx := entity.DB().Where("id = ?", purchase.PaymentMethodID).First(&paymentMethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paymentMethod not found"})
		return
	}

	// ค้นหา receipt ด้วย id
	if tx := entity.DB().Where("id = ?", purchase.ReceiptID).First(&receipt); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receipt not found"})
		return
	}

	// ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", product.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// ค้นหา product ด้วย id
	if tx := entity.DB().Where("id = ?", purchase.ProductID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	// สร้าง Purchase
	cpc := entity.Purchase{
		Time:          purchase.Time,
		PaymentMethod: paymentMethod,
		Receipt:       receipt,
		User:          user,
		Product:       product,
	}

	// บันทึก
	if err := entity.DB().Create(&cpc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": cpc})
}

// GET /purchase/:id
func GetPurchase(c *gin.Context) {
	var purchase entity.Purchase
	id := c.Param("id")

	if tx := entity.DB().Preload("PaymentMethod").Preload("Receipt").Preload("User").Preload("Product").Where("id = ?", id).First(&purchase); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "purchase not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": purchase})
}

// GET /purchases
func ListPurchases(c *gin.Context) {
	var purchases []entity.Purchase
	if err := entity.DB().Preload("PaymentMethod").Preload("Receipt").Preload("User").Preload("Product").Raw("SELECT * FROM products").Find(&purchases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": purchases})
}

// DELETE /purchase/:id
func DeletePurchase(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM purchases WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "purchase not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /purchase
func UpdatePurchase(c *gin.Context) {
	var purchase entity.Purchase
	var paymentMethod entity.PaymentMethod
	var receipt entity.Receipt
	var user entity.User
	var product entity.Product

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา paymentMethod ด้วย id
	if tx := entity.DB().Where("id = ?", purchase.PaymentMethodID).First(&paymentMethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paymentMethod not found"})
		return
	}

	// ค้นหา receipt ด้วย id
	if tx := entity.DB().Where("id = ?", purchase.ReceiptID).First(&receipt); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receipt not found"})
		return
	}

	// ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", product.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// ค้นหา product ด้วย id
	if tx := entity.DB().Where("id = ?", purchase.ProductID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	// อัพเดท Purchase
	upc := entity.Purchase{
		Time:          purchase.Time,
		PaymentMethod: paymentMethod,
		Receipt:       receipt,
		User:          user,
		Product:       product,
	}

	// บันทึก
	if err := entity.DB().Where("id = ?", purchase.ID).Updates(&upc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": upc})

}
