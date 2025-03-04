package cli

import (
	"fmt"
	"os"
)

// GetArg returns the os.Arg at the index, or empty string if the arg does not
// exist or is empty
func GetArg(index int) string {
	if index >= len(os.Args) {
		return ""
	}
	return os.Args[index]
}

// ExitOnEmpty prints the message to stderr and exits the program with status
// code 1 when any of the passed args is empty
func ExitOnEmpty(message string, args ...string) {
	for _, arg := range args {
		if arg != "" {
			continue
		}
		fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}
}

// ExitWithMessage prints the message to stderr and exists with status code 0
// when arg matches any of the patterns
func ExitWithMessage(message, arg string, patterns ...string) {
	for _, p := range patterns {
		if arg != p {
			continue
		}
		fmt.Fprintln(os.Stderr, message)
		os.Exit(0)
	}
}
