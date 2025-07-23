package email

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidEmail(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		cases := map[string]string{
			"simple":    "user@domain.com",
			"with_dot":  "user.name@domain.com",
			"with_tag":  "user+name@domain.com",
			"with_both": "user.name+tag@domain.com",
		}
		for name, testEmail := range cases {
			t.Run(name, func(t *testing.T) {
				require.True(t, IsValidEmail(testEmail))
			})
		}
	})
	t.Run("error", func(t *testing.T) {
		cases := map[string]string{
			"missing_at":          "userdomain.com",
			"double_at":           "user@domain@domain.com", // Detects an issue in IsValidEmail
			"no_name":             "@domain.com",            // Detects an issue in IsValidEmail
			"no_domain":           "user",
			"missing_domain_part": "user@domain",
			"empty_domain_part":   "user@domain.",
			"invalid_character":   "ąę@domain.com",
		}
		for name, testEmail := range cases {
			t.Run(name, func(t *testing.T) {
				require.False(t, IsValidEmail(testEmail))
			})
		}
	})
}
