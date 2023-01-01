package main

import (
"net/http"
"github.com/gin-gonic/gin"
)
type STKPushRequest struct {
	TargetNumber int `json:"targetNumber"`
	RequestID string `json:"requestID"`
}

func main() {


	// Create a new Gin router
	r := gin.New()

	// Define a route for the API
	r.POST("/pushstk", func(c *gin.Context) {
		// Bind the JSON data to a STKPushRequest struct
		var req STKPushRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		
		bot(req.TargetNumber, req.RequestID)

		// Return the BMI to the client
		c.JSON(http.StatusOK, gin.H{
			"message": "done",
		})
	})

	// Start the server
	r.Run()
}





func bot(targetNumber int, requestID string) {
	// get daraja brear token.
	ScammerStkPush(BearerTokenGenerator(), targetNumber, requestID)

}
