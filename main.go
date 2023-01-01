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
}
type STKCallback struct {
	Body struct {
		STKCallback struct {
			CallbackMetadata struct {
				Item []struct {
					Name  string `json:"Name"`
					Value string `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			MerchantRequestID string `json:"MerchantRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
		} `json:"stkCallback"`
	} `json:"Body"`
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
		bot(req.TargetNumber, req.RequestID)

		// Return a message to the client indicating that the request has been processed
		c.JSON(http.StatusOK, gin.H{
			"message": "done",
		})
	})

	// Define a route for handling callbacks from the STK API
	r.POST("/stkcallback/:param", func(c *gin.Context) {
		// Extract the "param" path parameter from the request
		param := c.Param("param")
		fmt.Println(param)
		// var data map[string]interface{}
		// // Parse and validate the request body as JSON
		// if err := c.ShouldBindJSON(&data); err != nil {
		var callbackData STKCallback
		if err := c.ShouldBindJSON(&callbackData); err != nil {
			// If the request body is invalid, return a Bad Request error
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(data)

	})
	// Start the server
	r.Run()
}

// bot is a function that handles STK push requests
func bot(targetNumber int, requestID string) {
	// Get a Daraja Bearer token
	ScammerStkPush(BearerTokenGenerator(), targetNumber, requestID)
}
