package main

import (
	"genai/genai/api"
	"genai/genai/db"
	"genai/genai/func/prompt"

	"github.com/gin-gonic/gin"
)


func main() {
	
	dbChat := []db.Message{}


	router := gin.Default()
	router.GET("/chat", func(c *gin.Context) {
		api.GetChat(c, dbChat)
	})
	router.POST("/prompt", func(c *gin.Context) {
		var req struct {
        Message string `json:"message"`
    }
    if err := c.BindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "invalid request"})
        return
    }
    prompt.Prompt(c, &dbChat, req.Message)
	})
	router.Run("localhost:8080")
	
}

