package user

import (
	"exercise4/email"
)

const emailSender = "no-reply@test.com"

type Service struct {
	Email email.Sender
}

func (s *Service) NotifyAccountCreated(email string) error {
	subject := "Account created"
	body := "Your account has been created."

	return s.Email.Send(emailSender, emailSender, subject, body)
}
