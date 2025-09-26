package generate

import (
	"context"
	"genai/genai/db"
	"genai/genai/pkg"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateChat(messages []db.Message) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(pkg.LoadEnv("GEMINI_API_KEY")))
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