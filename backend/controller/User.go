package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/flugika/seller-web-app/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// POST /user
func CreateUser(c *gin.Context) {
	var user entity.User
	var gender entity.Gender
	var district entity.District
	var occupation entity.Occupation
	var salary entity.Salary

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", user.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// ค้นหา district ด้วย id
	if tx := entity.DB().Where("id = ?", user.DistrictID).First(&district); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "district not found"})
		return
	}

	// ค้นหา occupation ด้วย id
	if tx := entity.DB().Where("id = ?", user.OccupationID).First(&occupation); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "occupation not found"})
		return
	}

	// ค้นหา salary ด้วย id
	if tx := entity.DB().Where("id = ?", user.SalaryID).First(&salary); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "salary not found"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	// สร้าง User
	cu := entity.User{
		Gender:     gender,
		District:   district,
		Occupation: occupation,
		Salary:     salary,
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   string(password),
		Role:       user.Role,
	}

	// บันทึก
	if err := entity.DB().Create(&cu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": cu})
}

// GET /user/:id
func GetUser(c *gin.Context) {
	var user entity.User
	id := c.Param("id")

	if tx := entity.DB().Preload("Gender").Preload("District").Preload("Occupation").Preload("Salary").Where("id = ?", id).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users
func ListUsers(c *gin.Context) {
	var users []entity.User
	if err := entity.DB().Preload("Gender").Preload("District").Preload("Occupation").Preload("Salary").Raw("SELECT * FROM users").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DELETE /user/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /user
func UpdateUser(c *gin.Context) {
	var user entity.User
	var gender entity.Gender
	var district entity.District
	var occupation entity.Occupation
	var salary entity.Salary

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", user.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// ค้นหา district ด้วย id
	if tx := entity.DB().Where("id = ?", user.DistrictID).First(&district); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "district not found"})
		return
	}

	// ค้นหา occupation ด้วย id
	if tx := entity.DB().Where("id = ?", user.OccupationID).First(&occupation); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "occupation not found"})
		return
	}

	// ค้นหา salary ด้วย id
	if tx := entity.DB().Where("id = ?", user.SalaryID).First(&salary); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "salary not found"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	// อัพเดท User
	uu := entity.User{
		Gender:     gender,
		District:   district,
		Occupation: occupation,
		Salary:     salary,
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   string(password),
		Role:       user.Role,
	}

	// บันทึก
	if err := entity.DB().Where("id = ?", user.ID).Updates(&uu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": uu})

}
