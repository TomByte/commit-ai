package git

import (
	"errors"
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
	output, err := exec.Command("git", "diff", "--cached", ".", ":!go.sum", ":!go.mod").Output()

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

func CheckDiff(diff string) error {
	if len(diff) == 0 {
		return errors.New("There is no diff to generate a commit for")
	}

	if len(diff) > 4000 {
		return errors.New("Diff is too large - max 4000 characters")
	}

	return nil
}
