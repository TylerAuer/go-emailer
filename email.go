package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

type Email struct {
	FromName  string
	FromEmail string
	ToName    string
	ToEmail   string
	Subject   string
	Message   string
}

func (em *Email) sendMailFromEmail() error {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	cl := getClient(config)

	gmailService, err := gmail.New(cl)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	from := mail.Address{em.FromName, em.FromEmail}
	to := mail.Address{em.ToName, em.ToEmail}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = em.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + em.Message

	gmsg := gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(msg)),
	}

	_, err = gmailService.Users.Messages.Send("me", &gmsg).Do()
	if err != nil {
		log.Printf("em %v, err %v", gmsg, err)
		return err
	}
	return err
}
