package api

import (
	"genai/genai/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetChat retrieves the chat history.
func GetChat(c *gin.Context, messages []db.Message) {
	c.IndentedJSON(http.StatusOK, messages)
}