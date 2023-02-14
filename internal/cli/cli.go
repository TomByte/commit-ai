package cli

import (
	"fmt"
	"strings"
)

func Confirm() bool {
	var input string

	fmt.Printf("Do you want to commit with this message? [y|n]: ")
	_, _ = fmt.Scanln(&input)

	input = strings.ToLower(input)

	if input == "y" || input == "yes" {
		return true
	}

	return false
}
