package main

import (
	"bufio"
	"fmt"
	"genai/genai/api"
	"genai/genai/db"
	"genai/genai/generate"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)


func main() {
	
	reader := bufio.NewReader(os.Stdin)
	var mu sync.Mutex
	dbChat := []db.Message{}


	go func() {
		router := gin.Default()
		router.GET("/chat", func(c *gin.Context) {
			mu.Lock()
			defer mu.Unlock()
			api.GetChat(c, dbChat)
		})
		router.Run("localhost:8080")
	}()
	
	for {
	fmt.Print("You: ")
		inputUser, _ := reader.ReadString('\n')
		inputUser = strings.TrimSpace(inputUser)

		if inputUser == "exit" {
			fmt.Println("Bye!")
			break
		}
	
	dbChat = append(dbChat, db.Message{Role: "user", Content: inputUser})

	newChat, err := generate.GenerateChat(dbChat)
	if err != nil {
		log.Fatal(err)
	}

	dbChat = append(dbChat, db.Message{Role: "assistant", Content: newChat})

	for _, m := range dbChat {
		fmt.Printf("[%s]: %s\n", m.Role, m.Content)
	}


	}
}

