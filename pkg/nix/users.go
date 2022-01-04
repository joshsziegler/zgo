// nix provides utilities for working with Linux and other Unix-like OSs.

package nix

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUsersFrom returns all users with a UID greater than 1000 from /etc/passwd
// like files. This ignores comments (lines starting with '#') and empty lines.
// All other lines MUST be colon-delimited and have seven parts.
func GetUsersFrom(path string) ([]string, error) {
	users := make([]string, 0)
	file, err := os.ReadFile(path)
	if err != nil {
		return users, fmt.Errorf("cannot read users from: %s; %w", path, err)
	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue // skip comments
		} else if len(strings.TrimSpace(line)) == 0 {
			continue // skip empty lines
		}
		parts := strings.Split(line, ":")
		if len(parts) != 7 { // we could simply check for the parts we need, but this is stronger
			return users, fmt.Errorf("users in /etc/passwd like files must have 7 parts: %s", line)
		}
		username := parts[0]
		uid, err := strconv.Atoi(parts[2])
		if err != nil { // provide path and line number on errors
			return users, fmt.Errorf("cannot convert UID '%s' to an integer from %s:%d", parts[2], path, i)
		}
		if uid >= 1000 { // Only return users who are not system or service accounts
			users = append(users, username)
		}
	}
	return users, nil
}
