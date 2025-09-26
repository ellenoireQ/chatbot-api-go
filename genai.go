package main

import (
	"bufio"
	"fmt"
	"genai/genai/db"
	"genai/genai/generate"
	"log"
	"os"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	dbChat := []db.Message{}

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

