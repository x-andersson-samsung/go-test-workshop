package email

import "time"

type Sender struct {
	// Connection information
}

func (s *Sender) Send(sender, recipient string, subject string, body string) error {
	// Simulate connecting to an SMTP server and sending data
	time.Sleep(1 * time.Second)
	return nil
}
