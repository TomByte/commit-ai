package git

import (
	"commit-ai/internal/log"
	"os/exec"
)

type GIT struct {
}

func NewGIT() *GIT {
	return &GIT{}
}

func (g *GIT) IsRepo() bool {
	cmd, err := exec.Command("git", "status").Output()

	if err != nil {
		log.Error(string(cmd))
		return false
	}

	log.Info(string(cmd))
	return true
}

func (g *GIT) GetDiff() (string, error) {

	output, err := exec.Command("git", "diff", "--cached", ".").Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}
