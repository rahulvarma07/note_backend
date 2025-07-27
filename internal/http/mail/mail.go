package mail

import (
	"log"

	"github.com/rahulvarma07/note_backend/internal/config"
	"gopkg.in/gomail.v2"
)

func SendMail(mailConfig *config.Mail, userEmail string) {
	// what all is needed
	// myMail userMail details of myMail port pass etc..

	from := mailConfig.SenderMailID // app mail id
	to := userEmail                 // user mail id
	mailHost := mailConfig.MailHost
	mailPort := mailConfig.MailPort         // 587
	mailPassword := mailConfig.MailPassword // mail app password

	htmlBody := 
	`<html>
		<body style="font-family: Arial, sans-serif; line-height: 1.6;">
			<h2>Verify Your Email Address</h2>
			<p>Hello,</p>
			<p>Thank you for signing up! Please click the button below to verify your email address:</p>
			<a href=""
				style="display: inline-block; 
					   padding: 12px 24px; 
					   font-size: 16px; 
					   color: white; 
					   background-color: #28a745; 
					   text-decoration: none; 
					   border-radius: 6px;
					   margin: 20px 0;
					   ">
				Verify Email
			</a>
			<p>If you did not request this, you can ignore this email.</p>
			<p>Thanks,<br>The YourApp Team</p>
		</body>
	</html>`

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Verify Your Email")
	mailer.SetBody("text/Html", htmlBody)

	dialer := gomail.NewDialer(mailHost, mailPort, from, mailPassword)

	if err := dialer.DialAndSend(); err != nil{
		log.Fatal("unable to send the verifiaction mail")
	}

	log.Println("Successfully send the verification mail")
}
