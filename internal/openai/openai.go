package openai

import (
	"commit-ai/internal/config"
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

const (
	Prompt = "I want you to act like a git commit message writer. I will input a git diff and your job is to convert it into a useful commit message. Do not preface the commit with anything, use the present tense, return a complete sentence, and do not repeat yourself: \n"
)

type OpenAI struct {
	Key    string
	Url    string
	Client *resty.Client
}

type CompletionsRequest struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        float64 `json:"max_tokens"`
	Temperature      float64 `json:"temperature"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	Stream           bool    `json:"stream"`
	N                float64 `json:"n"`
}

type CompletionsResponse struct {
	Choices []Choice `json:"choices"`
}

type CompletionsErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
}

type Choice struct {
	Text string `json:"text"`
}

func NewOpenAI(config *config.Config) *OpenAI {
	return &OpenAI{
		Key:    config.OpenAIKey,
		Url:    config.OpenAIUrl,
		Client: resty.New(),
	}
}

func (ai *OpenAI) GetCommitMessage(diff string) (string, error) {
	request := CompletionsRequest{
		Model:            "text-davinci-003",
		Prompt:           fmt.Sprintf("%s%s", Prompt, diff),
		MaxTokens:        200,
		Temperature:      0.7,
		TopP:             1,
		FrequencyPenalty: 1,
		PresencePenalty:  1,
		Stream:           false,
		N:                1,
	}

	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Start()
	s.Suffix = " Generating commit"

	resp, err := ai.Client.R().
		SetHeader("Authorization", "Bearer "+ai.Key).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetError(&CompletionsErrorResponse{}).
		SetResult(&CompletionsResponse{}).
		Post(fmt.Sprintf(ai.Url + "/completions"))

	s.Stop()

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", errors.New(resp.Result().(*CompletionsErrorResponse).Error.Message)
	}

	result := resp.Result().(*CompletionsResponse)

	return formatCommit(result.Choices[0].Text), nil
}

func formatCommit(commit string) string {
	c := strings.ReplaceAll(commit, "\n", "")
	return strings.TrimRight(c, ".")
}
