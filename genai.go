package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Message struct {
	Role    string
	Content string
}

func loadEnv(key string) string{
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
func generateChat(messages []Message) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(loadEnv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-pro")

	var userInput string
	if len(messages) > 0 {
		userInput = messages[len(messages) -1].Content
	} else {
		userInput = "Hello"
	}

	resp, err := model.GenerateContent(ctx, genai.Text(userInput))
	if err != nil {
		return "", err
	}

	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			if text, ok := part.(genai.Text); ok {
				return string(text), nil
			}
		}
	}
	return "", nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	dbChat := []Message{}

	for {
	fmt.Print("You: ")
		inputUser, _ := reader.ReadString('\n')
		inputUser = strings.TrimSpace(inputUser)

		if inputUser == "exit" {
			fmt.Println("Bye!")
			break
		}
	
	dbChat = append(dbChat, Message{Role: "user", Content: inputUser})

	newChat, err := generateChat(dbChat)
	if err != nil {
		log.Fatal(err)
	}

	dbChat = append(dbChat, Message{Role: "assistant", Content: newChat})

	for _, m := range dbChat {
		fmt.Printf("[%s]: %s\n", m.Role, m.Content)
	}
	}
}

