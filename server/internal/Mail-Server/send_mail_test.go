package mail

import (
	"context"
	"testing"
)

func (ms *MailService) TestSendMail(t *testing.T) {

  err := ms.SendMail(context.Background() , SendMailContent{
    senderEmail: ms.TestEmail,
    emailSubject: "testing send email",
    emailBody: "test successfull!",
  })

  if err != nil {
    t.Fatalf("Error: test failed in Send Mail with error: %v" , err.Error())
  }
}
