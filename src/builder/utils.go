package builder

import (
	"fmt"
	"os"
	"strings"
)

// fatalError prints an error message and exits the program
func fatalError(context string, err error) {
	fmt.Printf("Error [%s]: %v\n", context, err)
	os.Exit(1)
}

// replaceLastOccurrence replaces the last occurrence of a substring
func replaceLastOccurrence(s, old, new string) string {
	lastIndex := strings.LastIndex(s, old)
	if lastIndex == -1 {
		return s
	}
	return s[:lastIndex] + new + s[lastIndex+len(old):]
}
