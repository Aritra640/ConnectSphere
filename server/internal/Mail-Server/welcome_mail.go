package mail

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

//A welcome mail send to greet the user with username and email
func (ms *MailService) WelcomeMail(ctx context.Context, username string, email string, wg *sync.WaitGroup) error {
	defer wg.Done()

	bodyStr := fmt.Sprintf(`
<div style="font-family: Arial, sans-serif; max-width: 600px; padding: 20px; border: 1px solid #ddd; border-radius: 10px; background-color: #f9f9f9;">
    <h1 style="color: #333; text-align: center;">Hello %s,</h1>

    <p style="font-size: 16px; color: #555; text-align: center;">
        Welcome to our app! We're excited to have you on board.
    </p>

    <p style="font-size: 16px; color: #555; text-align: center;">
        You signed up with the email: <strong style="color: #333;">%s</strong>
    </p>

    <p style="font-size: 14px; color: #777; text-align: center;">
        If this wasn't you, please contact our support team immediately.
    </p>

    <div style="text-align: center; margin: 30px 0;">
        <a href="https://yourapp.example.com" style="display: inline-block; background-color: #007bff; color: white; text-decoration: none; padding: 12px 24px; border-radius: 5px; font-size: 16px;">
            Go to Dashboard
        </a>
    </div>

    <hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">

    <p style="text-align: center; font-size: 12px; color: #aaa;">
        &copy; 2025 Your Company. All rights reserved.
    </p>
</div>
`, username, email)

	errChan := make(chan error)
	defer close(errChan)

	go func() {

		err := ms.SendMail(ctx, SendMailContent{
			emailBody:    bodyStr,
			emailSubject: "Welcome to ConnectSphere",
			senderEmail:  email,
		})

		errChan <- err
	}()

	select {
	case <-ctx.Done():
		log.Println("Timed out in WelcomeMail")
		return errors.New("Timed out")

	case err := <-errChan:
		if err == nil {
			return nil
		}
		log.Println("Failed to send Welcome mail")
		return err
	}
}


