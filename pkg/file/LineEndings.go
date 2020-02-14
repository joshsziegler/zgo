package file

import "strings"

// RemoveCR returns a new string with all carriage returns removed (Unix Style).
//
// This converts CR line endings to to LineFeeds (i.e. \n).
func RemoveCR(text string) string {
	// Replace all DOS-style CR+LF to LF
	text = strings.Replace(text, "\r\n", "\n", -1)
	// Replace all Classic Mac OS-style CR to LF
	text = strings.Replace(text, "\r", "\n", -1)
	return text
}
