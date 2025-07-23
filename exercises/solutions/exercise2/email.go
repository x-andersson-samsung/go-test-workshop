package email

import (
	"strings"
)

const allowedCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&*+-/=?^_.@"

func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}

	// @ check
	// Fix #1 - should not allow more than 1 @ character
	if strings.Count(email, "@") != 1 {
		return false
	}

	// Character check
	for _, c := range email {
		if !strings.Contains(allowedCharacters, string(c)) {
			return false
		}
	}

	// Domain checks - no need to check parts length. Checks above guarantee presence of @ and at least 2 parts
	parts := strings.Split(email, "@")

	// Fix #2 - should not allow empty username
	if len(parts[0]) == 0 {
		return false
	}

	if parts[1] == "" {
		// No domain
		return false
	}

	domainParts := strings.Split(parts[len(parts)-1], ".")
	if len(domainParts) < 2 {
		return false
	}

	for _, p := range domainParts {
		if len(p) == 0 {
			return false
		}
	}

	return true
}
