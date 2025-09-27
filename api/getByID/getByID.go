package getByID

import (
	"genai/genai/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetByID(c *gin.Context, id int, message []db.Message) {
	for _, msg := range message {
		if msg.ID == id {
			c.IndentedJSON(http.StatusOK, msg)
			return
		}
	
		}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "ID not found"})
}