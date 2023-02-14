package openai

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

const (
	Prompt = "I want you to act like a git commit message writer. I will input a git diff and your job is to convert it into a useful commit message. Do not preface the commit with anything, use the present tense, return a complete sentence, and do not repeat yourself: "
)

type OpenAI struct {
	Key    string
	Url    string
	Client *resty.Client
}

type CompletionsRequest struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature int    `json:"temperature"`
}

type CompletionsResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text string `json:"text"`
}

func NewOpenAI() *OpenAI {
	return &OpenAI{
		Key:    os.Getenv("OPENAI_KEY"),
		Url:    "https://api.openai.com/v1",
		Client: resty.New(),
	}
}

func (ai *OpenAI) GetCommitMessage(diff string) (string, error) {
	request := CompletionsRequest{
		Model:       "text-davinci-003",
		Prompt:      fmt.Sprintf("%s%s", Prompt, diff),
		MaxTokens:   4000,
		Temperature: 0,
	}

	body, _ := json.Marshal(request)

	resp, err := ai.Client.R().
		SetHeader("Authorization", "Bearer "+ai.Key).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&CompletionsResponse{}).
		Post(fmt.Sprintf(ai.Url + "/completions"))

	if err != nil {
		return "", err
	}

	result := resp.Result().(*CompletionsResponse)

	return result.Choices[0].Text, nil
}
