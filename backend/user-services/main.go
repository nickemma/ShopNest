package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Connect to PostgreSQL
	dsn := "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{})

	// HTTP Server
	r := gin.Default()
	r.GET("/api/users/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User Service: Hello!"})
	})

	r.POST("/api/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	})

	r.Run(":8080")
}
