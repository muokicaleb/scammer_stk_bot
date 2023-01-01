package main

import (
"net/http"
"github.com/gin-gonic/gin"
"encoding/json"
"fmt"

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

	r.POST("/stkcallback/:param", func(c *gin.Context) {
        param := c.Param("param")
		 // Create a map to store the data
		 var data map[string]interface{}

		 // Unmarshal the JSON response into the map
		 if err := json.Unmarshal([]byte(param), &data); err != nil {
			fmt.Println(err)
		 }
	 
		 // Print the map
		 fmt.Printf("%+v\n", data)
	
	})
	// Start the server
	r.Run()
}





func bot(targetNumber int, requestID string) {
	// get daraja brear token.
	ScammerStkPush(BearerTokenGenerator(), targetNumber, requestID)

}
