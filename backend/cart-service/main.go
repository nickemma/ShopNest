package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DynamoDB client (local for dev)
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://dynamodb:8000"), // Local DynamoDB
	}))
	db := dynamodb.New(sess)

	// HTTP Server
	r := gin.Default()
	r.GET("/api/carts/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Cart Service: Hello!"})
	})

	r.POST("/api/carts", func(c *gin.Context) {
		// Example: Add item to cart
		_, err := db.PutItem(&dynamodb.PutItemInput{
			TableName: aws.String("carts"),
			Item: map[string]*dynamodb.AttributeValue{
				"userId": {S: aws.String("123")},
				"items":  {L: []*dynamodb.AttributeValue{}},
			},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Cart created"})
	})

	r.Run(":8080")
}
