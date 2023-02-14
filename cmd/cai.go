package cmd

import (
	"commit-ai/internal/cli"
	"commit-ai/internal/git"
	"commit-ai/internal/log"
	"commit-ai/internal/openai"
)

func Exec() {
	g := git.NewGIT()
	ai := openai.NewOpenAI()

	if !g.IsRepo() {
		log.Error("Current directory is not a GIT repository")
		return
	}

	diff, err := g.GetDiff()
	if err != nil {
		log.Error("Unable to create a usable diff for message")
		return
	}

	msg, err := ai.GetCommitMessage(diff)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Generated Commit Message: ", msg)

	if !cli.Confirm() {
		return
	}

	err = g.Commit(msg)

	if err != nil {
		log.Error(err)
	}
}
