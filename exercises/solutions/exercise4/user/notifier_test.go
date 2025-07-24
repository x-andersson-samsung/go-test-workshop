package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// Manual mock example
type mockSender struct {
	SendFn func(sender string, email string, subject string, body string) error
}

func (m *mockSender) Send(sender string, email string, subject string, body string) error {
	return m.SendFn(sender, email, subject, body)
}

func TestService_NotifyAccountCreated(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		recipient := "test@test.com"
		notifier := Service{
			Email: &mockSender{
				SendFn: func(sender string, email string, subject string, body string) error {
					require.Equal(t, emailSender, sender)
					require.Equal(t, recipient, email)
					require.Equal(t, accountCreatedSubject, subject)
					require.Equal(t, accountCreatedBody, body)
					return nil
				},
			},
		}

		require.NoError(t, notifier.NotifyAccountCreated(recipient))
	})
	t.Run("error", func(t *testing.T) {
		expected := errors.New("test error")
		notifier := Service{
			Email: &mockSender{
				SendFn: func(sender string, email string, subject string, body string) error {
					return expected
				},
			},
		}

		require.ErrorIs(t, notifier.NotifyAccountCreated("any"), expected)
	})
}
