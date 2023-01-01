package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// STKPushRequest represents the data in the POST request for the /pushstk route
type STKPushRequest struct {
	TargetNumber int    `json:"targetNumber"`
	RequestID    string `json:"requestID"`
	PushAmount   int    `json:"pushAmount"`
}

func main() {

	// Create a new Gin router
	r := gin.New()

	// Define a route for the API
	r.POST("/pushstk", func(c *gin.Context) {
		// Bind the JSON data in the request body to a STKPushRequest struct
		var req STKPushRequest
		if err := c.BindJSON(&req); err != nil {
			// If the request body is invalid, return a Bad Request error
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Call the bot function with the TargetNumber and RequestID from the request
		bot(req.TargetNumber, req.RequestID, req.PushAmount)

		// Return a message to the client indicating that the request has been processed
		c.JSON(http.StatusOK, gin.H{
			"message": "stk push sent",
		})
	})

	// Define a route for handling callbacks from the STK API
	r.POST("/stkcallback/:param", func(c *gin.Context) {
		// Extract the "param" path parameter from the request
		param := c.Param("param")
		fmt.Println(param)
		var data map[string]interface{}

		// Parse and validate the request body as JSON

		if err := c.ShouldBindJSON(&data); err != nil {
			// If the request body is invalid, return a Bad Request error
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		str := fmt.Sprintf("%+v", data)
		fmt.Println(str)

	})
	// Start the server
	r.Run()
}

// bot is a function that handles STK push requests
func bot(targetNumber int, requestID string, pushAmount int) {
	// Get a Daraja Bearer token
	ScammerStkPush(BearerTokenGenerator(), targetNumber, requestID, pushAmount)
}
