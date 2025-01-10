package services

type SMSService struct {
}

func (s *SMSService) Send(msg string) bool {
	return true
}

type EmailService struct {
}

func (s *EmailService) Send(msg string) bool {
	return true
}
