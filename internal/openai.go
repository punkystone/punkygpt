package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type OpenAIRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type OpenAIResponse struct {
	Output []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Status  string `json:"status"`
		Content []struct {
			Type        string `json:"type"`
			Annotations []any  `json:"annotations"`
			Logprobs    []any  `json:"logprobs"`
			Text        string `json:"text"`
		} `json:"content"`
		Role string `json:"role"`
	} `json:"output"`
}

func GetResponse(user string, message string, apiKey string) (string, error) {
	prompt, err := os.ReadFile("prompt.txt")
	if err != nil {
		return "", err
	}
	input := fmt.Sprintf(string(prompt), user, message)
	requestJSON := OpenAIRequest{
		Model: "gpt-4.1",
		Input: input}
	buffer := &bytes.Buffer{}
	err = json.NewEncoder(buffer).Encode(requestJSON)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/responses", buffer)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	openAIResponse := &OpenAIResponse{}
	err = json.NewDecoder(response.Body).Decode(openAIResponse)
	if err != nil {
		return "", err
	}
	if len(openAIResponse.Output) == 0 {
		return "", errors.New("no output from openai")
	}
	if len(openAIResponse.Output[0].Content) == 0 {
		return "", errors.New("no content in output from openai")
	}
	return openAIResponse.Output[0].Content[0].Text, nil
}
