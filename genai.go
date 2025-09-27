package main

import (
	GetByID "genai/genai/api/GetByID"
	GetChat "genai/genai/api/GetChat"
	"genai/genai/db"
	"genai/genai/func/prompt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)


func main() {
	
	dbChat := []db.Message{}


	router := gin.Default()
	router.GET("/chat", func(c *gin.Context) {
		GetChat.GetChat(c, dbChat)
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

	router.GET("/history/:id", func(c *gin.Context) {
		id := c.Param("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Invalid ID:", err)
		}
		GetByID.GetByID(c, num, dbChat)
	})
	router.Run("localhost:8080")
	
}

