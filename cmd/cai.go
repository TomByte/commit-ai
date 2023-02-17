package cmd

import (
	"commit-ai/internal/cli"
	"commit-ai/internal/config"
	"commit-ai/internal/git"
	"commit-ai/internal/log"
	"commit-ai/internal/openai"
	"fmt"
)

func Exec() {
	c := config.NewConfig()
	err := c.Load()

	if err != nil {
		log.Error(err)
		return
	}

	g := git.NewGIT()
	ai := openai.NewOpenAI(c)

	if !g.IsRepo() {
		log.Error("Current directory is not a GIT repository")
		return
	}

	diff, err := g.GetDiff()

	if err != nil {
		log.Error("Unable to create a usable diff for message")
		return
	}

	err = git.CheckDiff(diff)

	if err != nil {
		log.Error(err)
		return
	}

	msg, err := ai.GetCommitMessage(diff)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info(fmt.Sprintf("Generated commit message: %s", msg))

	if !cli.Confirm() {
		return
	}

	err = g.Commit(msg)

	if err != nil {
		log.Error(err)
	}
}
