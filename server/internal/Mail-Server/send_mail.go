package mail

import (
	"context"
	"log"

	"github.com/wneessen/go-mail"
)

type SendMailContent struct {
	senderEmail  string
	emailSubject string
	emailBody    string
}

// Send mail with a sender email, email body, email subject
func (ms *MailService) SendMail(ctx context.Context, content SendMailContent) error {

	message := mail.NewMsg()

	if err := message.From(ms.Email); err != nil {
		log.Println("Error: Failed to set FROM address in SendMail: ", err)
		return err
	}

	if err := message.To(content.senderEmail); err != nil {
		log.Println("Error: Failed to set TO address in SendMail: ", err)
		return err
	}

	message.Subject(content.emailSubject)
	message.SetBodyString(mail.TypeTextHTML, content.emailSubject)

	client, err := mail.NewClient(
		ms.Smtp,
		mail.WithPort(587),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(ms.Email),
		mail.WithPassword(ms.Password),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		log.Println("Error: Failed to create mail client: ", err)
		return err
	}

	if err = client.DialAndSend(message); err != nil {
		log.Println("Error: Failed to send mail: ", err)
		return err
	}

	log.Println("one mail successfully send to mailid: ", content.senderEmail)
	return nil
}
