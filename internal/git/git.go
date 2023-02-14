package git

import (
	"os/exec"
)

type GIT struct {
}

func NewGIT() *GIT {
	return &GIT{}
}

func (g *GIT) IsRepo() bool {
	cmd := exec.Command("git", "status")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func (g *GIT) GetDiff() (string, error) {
	output, err := exec.Command("git", "diff", "--cached", ".").Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (g *GIT) Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
