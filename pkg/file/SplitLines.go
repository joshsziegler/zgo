package file

import "strings"

// SplitLines returns the string split by new lines (CR, CR+LF, and LF).
func SplitLines(input string) []string {
	// Convert Carriage Returns to Line Feeds (i.e. Unix-style line endings)
	input = RemoveCR(input)
	// Split by Line Feeds
	lines := strings.Split(input, "\n")
	return lines
}
