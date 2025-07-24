package user

const (
	emailSender           = "no-reply@test.com"
	accountCreatedSubject = "Account Created"
	accountCreatedBody    = "Your account has been created."
)

// Create the interface
type Sender interface {
	Send(sender, recipient, subject, body string) error
}

type Service struct {
	// Change field type from concrete struct to the interface
	Email Sender
}

func (s *Service) NotifyAccountCreated(email string) error {
	// Fix second parameter to properly pass the recipient email
	return s.Email.Send(emailSender, email, accountCreatedSubject, accountCreatedBody)
}
