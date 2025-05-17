package mail

import (
	"context"
	"fmt"
	"log"
)

func (ms *MailService) OTPmail(ctx context.Context, username string, email string, otp string) error {

	bodyStr := fmt.Sprintf(`
    <div style="font-family: Arial, sans-serif; max-width: 600px; padding: 20px; border: 1px solid #ddd; border-radius: 10px; background-color: #f9f9f9;">
        <h1 style="color: #333; text-align: center;">Hello %s,</h1>
        <p style="font-size: 16px; color: #555; text-align: center;">
            Your registered email is <strong style="color: #333;">%s</strong>.
        </p>
        <p style="font-size: 18px; text-align: center; font-weight: bold; color: #2d89ff;">
            Your OTP is: <span style="font-size: 22px; background-color: #e0e0e0; padding: 5px 10px; border-radius: 5px;">%s</span>
        </p>
        <p style="font-size: 14px; color: #777; text-align: center;">
            Please do not share this OTP with anyone. It is valid for a limited time.
        </p>
        <hr style="border: none; border-top: 1px solid #ddd; margin: 20px 0;">
        <p style="text-align: center; font-size: 12px; color: #aaa;">
            &copy; 2025 Your Company. All rights reserved.
        </p>
    </div>
`, username, email, otp)

	errChan := make(chan error)
	defer close(errChan)
	go func() {

		err := ms.SendMail(ctx, SendMailContent{
			emailSubject: "ConnectSphere Email Verification",
			emailBody:    bodyStr,
			senderEmail:  email,
		})

		errChan <- err
	}()

	err := <-errChan

	if err != nil {
		log.Println("Failed to send otp mail : ", err)
		return err
	}
	return nil
}
