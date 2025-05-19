package mail

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

func (ms *MailService) ForgetPasswordMail(
	ctx context.Context, username string, email string, 
	resetLink string, wg *sync.WaitGroup , res chan error) {
	defer wg.Done()

	bodyStr := fmt.Sprintf(`
    <div style="font-family: Arial, sans-serif; max-width: 600px; padding: 20px; border: 1px solid #ddd; border-radius: 10px; background-color: #f9f9f9;">
        <h1 style="color: #333; text-align: center;">Password Reset Request</h1>
        <p style="font-size: 16px; color: #555; text-align: center;">
            Hi <strong style="color: #333;">%s</strong>, we received a request to reset your password for the account registered with <strong>%s</strong>.
        </p>
        <p style="font-size: 16px; color: #555; text-align: center;">
            To reset your password, click the button below:
        </p>
        <div style="text-align: center; margin: 20px 0;">
            <a href="%s" style="background-color: #2d89ff; color: white; padding: 12px 20px; text-decoration: none; border-radius: 5px; font-size: 16px;">
                Reset Password
            </a>
        </div>
        <p style="font-size: 14px; color: #777; text-align: center;">
            If you did not request a password reset, please ignore this email. This link will expire in 30 minutes.
        </p>
        <hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">
        <p style="text-align: center; font-size: 12px; color: #aaa;">
            &copy; 2025 Your Company. All rights reserved.
        </p>
    </div>
`, username, email, resetLink)

	select{
	case <-ctx.Done():
		log.Println("Timed out in forget mail sender")
		res <- errors.New("Timed out")
		return
	default: 
		//So context not canceled : continue 
	}

	log.Println("Continuing with sending the email in ForgetPasswordMail")

	errChan := make(chan error)
	defer close(errChan)
	go func() {
		err := ms.SendMail(ctx, SendMailContent{
			senderEmail:  email,
			emailSubject: "Forget Password for ConnectSphere",
			emailBody:    bodyStr,
		})

		errChan <- err
	}()

	err := <-errChan

	res <- err
}
