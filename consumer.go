package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for Recipient := range ch {

		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", Recipient.Email, "This is a test email")

		// msg := []byte(formattedMsg)
		msg, err := executeTemplate(Recipient)

		if err != nil {
			fmt.Println("worker %d error while parsing for email: %s", id, Recipient.Email)
			continue
		}

		fmt.Printf("Worker %d sending email to %s\n", id, Recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "mksop321@mail.com", []string{Recipient.Email}, []byte(msg))

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(50 * time.Millisecond)

		fmt.Println(id, Recipient)
		fmt.Printf("Worker %d sending email to %s\n", id, Recipient.Email)

	}

}
