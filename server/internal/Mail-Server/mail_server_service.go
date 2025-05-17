package mail

type MailService struct {
	Email     string
	Password  string
	Smtp      string
	TestEmail string
}


var MailSetup = &MailService{}
