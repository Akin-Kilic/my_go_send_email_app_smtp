package main

import (
	"log"

	"net/smtp"
)

func main() {

	// Choose auth method and set it up

	auth := smtp.PlainAuth("", "akin@gmail.com", "secret key", "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it

	to := []string{"akin@gmail.com"}

	msg := []byte("To: akin@gmail.com\r\n" +

		"Subject: Why aren’t you using Mailtrap yet?\r\n" +

		"\r\n" +

		"Here’s the space for our great sales pitch\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "akin@gmail.com", to, msg)

	if err != nil {

		log.Fatal(err)

	}

}
