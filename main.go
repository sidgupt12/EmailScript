package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// Sender data
	from := "sidgupt12@gmail.com"

	// Receiver email address
	to := []string{"smurfactval@gmail.com", "sidgupt12@gmail.com", "siddhant2213031@akgec.ac.in"}

	// SMTP server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, Password, smtpHost)

	// Sending email to each recipient.
	for _, recipient := range to {
		message := []byte("Subject: Test Email from Golang\n" +
			"To: " + recipient + "\n" +
			"From: sidgupt12@gmail.com\n\n" +
			"This is a test email sent from a Go script.")

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipient}, message)
		if err != nil {
			fmt.Println("Error sending email to", recipient, ":", err)
		} else {
			fmt.Println("Email sent successfully to", recipient)
		}
	}
}
