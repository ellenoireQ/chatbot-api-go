package prompt

import (
	"genai/genai/db"
	"genai/genai/func/generate"

	"github.com/gin-gonic/gin"
)

func Prompt(c *gin.Context, dbChat *[]db.Message, message string) {
    newChat, err := generate.GenerateChat(message)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    *dbChat = append(*dbChat, db.Message{
        User:      db.User{Role: "user", Content: message},
        Assistant: db.Assistant{Role: "assistant", Content: newChat},
    })

    c.IndentedJSON(200, *dbChat)
}
