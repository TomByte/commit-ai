package cmd

import (
	"commit-ai/internal/git"
	"commit-ai/internal/log"
	"commit-ai/internal/openai"
)

func Exec() {
	g := git.NewGIT()
	ai := openai.NewOpenAI()

	if !g.IsRepo() {
		log.Error("current directory is not a GIT repository")
		return
	}

	diff, err := g.GetDiff()
	if err != nil {
		log.Error("unable to create a usable diff for message")
		return
	}

	msg, err := ai.GetCommitMessage(diff)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info(msg)
}
