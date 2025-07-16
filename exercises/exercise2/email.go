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
	if strings.Count(email, "@") < 1 {
		return false
	}

	// Character check
	for _, c := range email {
		if !strings.Contains(allowedCharacters, string(c)) {
			return false
		}
	}

	// Domain checks
	parts := strings.Split(email, "@")
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
