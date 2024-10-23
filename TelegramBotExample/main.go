package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ChatGPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Panic("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	openAIToken := os.Getenv("OPENAI_API_KEY")
	if openAIToken == "" {
		log.Panic("OPENAI_API_KEY environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		log.Printf("[User: %s] %s", update.Message.From.UserName, update.Message.Text)

		responseText, err := getChatGPTResponse(update.Message.Text, openAIToken)
		if err != nil || responseText == "" {
			log.Println(err)
			responseText = "Sorry, there was an error processing your request or the response was empty."
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}

func getChatGPTResponse(userInput, token string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	requestBody := ChatGPTRequest{
		Model: "gpt-4o-mini",
		Messages: []Message{
			{Role: "user", Content: userInput},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received error: %d", resp.StatusCode)
		return "", nil
	}

	var chatGPTResponse ChatGPTResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatGPTResponse); err != nil {
		return "", err
	}

	if len(chatGPTResponse.Choices) > 0 {
		return chatGPTResponse.Choices[0].Message.Content, nil
	}

	return "", nil
}

// Ensure that you have your bot token set in an environment variable named TELEGRAM_BOT_TOKEN.
// You can set it like this:
// export TELEGRAM_BOT_TOKEN="your_bot_token"
// Also, set your OpenAI API key in an environment variable named OPENAI_API_KEY.
// export OPENAI_API_KEY="your_openai_api_key"
// Then run the program.
