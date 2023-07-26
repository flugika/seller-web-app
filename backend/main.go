package main

import (
	"os"

	"github.com/flugika/seller-web-app/controller"
	"github.com/flugika/seller-web-app/entity"
	"github.com/flugika/seller-web-app/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// Delete database file before BUILD and RUN
	os.Remove("./seller.db")

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// user Routes
			router.POST("/user", controller.CreateUser)
			router.GET("/user/:id", controller.GetUser)
			router.GET("/users", controller.ListUsers)
			router.DELETE("/user/:id", controller.DeleteUser)
			router.PATCH("/users", controller.UpdateUser)
		}
	}
	// login User Route
	r.POST("/login", controller.Login)

	r.POST("/signup", controller.CreateUser)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}
