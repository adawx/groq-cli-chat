package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Request Types

type InputMessage struct {
	Message string `json:"message"`
}

type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ApiReqBody struct {
	Messages []GroqMessage `json:"messages"`
	Model    string        `json:"model"`
}

// Response Types

type ChatCompletion struct {
	ID                string            `json:"id"`
	Object            string            `json:"object"`
	Created           int64             `json:"created"`
	Model             string            `json:"model"`
	Choices           []Choice          `json:"choices"`
	Usage             Usage             `json:"usage"`
	SystemFingerprint string            `json:"system_fingerprint"`
	XGroq             map[string]string `json:"x_groq"`
}

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int     `json:"prompt_tokens"`
	PromptTime       float64 `json:"prompt_time"`
	CompletionTokens int     `json:"completion_tokens"`
	CompletionTime   float64 `json:"completion_time"`
	TotalTokens      int     `json:"total_tokens"`
	TotalTime        float64 `json:"total_time"`
}

func main() {
	api_key := os.Getenv("GROQ_API_KEY")

	if api_key == "" {
		fmt.Println("GROQ_API_KEY environent variable is not set.")
		return
	}

	var message string
	var model string
	flag.StringVar(&message, "message", "", "Message to be sent to Groq API. String needs to be wrapped in double quotes.")
	flag.StringVar(&model, "model", "", "Model to be used by Groq API. Available models:  gemma-7b-it, llama3-8b-8192, llama3-70b-8192, mixtral-8x7b-32768.")
	flag.Parse()

	if message != "" && model != "" {
		sendPostRequest(api_key, message, model)
		return
	} else {
		fmt.Printf("Please provide a command.\nUsage: gchat -model <model> -message \"<message>\"\n")
		return
	}
}

// Function to send a post request to groq api
func sendPostRequest(api_key string, message string, model string) {
	url := "https://api.groq.com/openai/v1/chat/completions"
	jsonBody := ApiReqBody{
		Messages: []GroqMessage{
			{
				Role:    "user",
				Content: message,
			},
		},
		Model: model,
	}

	reqBody, err := json.Marshal(jsonBody)

	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var chatCompletion ChatCompletion
	unmarshalErr := json.Unmarshal([]byte(body), &chatCompletion)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return
	}

	fmt.Println(chatCompletion.Choices[0].Message.Content)
	return

}
