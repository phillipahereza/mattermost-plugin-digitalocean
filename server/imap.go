package main

import (
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

func (p *Plugin) connectToImapMailClient() *client.Client {
	config := p.getConfiguration()
	imapserver := config.IMAPServer
	imapusername := config.IMAPUsername
	imappassword := config.IMAPPassword

	c, err := client.DialTLS(imapserver, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Login
	if err := c.Login(imapusername, imappassword); err != nil {
		log.Fatal(err)
	}
	return c
}

func (p *Plugin) retrieveDOEmailAlerts() ([]string, error) {
	var alerts []string
	mailClient := p.connectToImapMailClient()

	defer mailClient.Logout()

	_, err := mailClient.Select("INBOX", false)
	if err != nil {
		return []string{}, err
	}

	criteria := imap.NewSearchCriteria()

	//Select only unseen messages
	criteria.WithoutFlags = []string{"\\Seen"}

	// Only Digital Ocean Monitoring alerts
	criteria.Header.Add("SUBJECT", "DigitalOcean monitoring triggered")

	ids, _ := mailClient.Search(criteria)

	seqset := new(imap.SeqSet)
	seqset.AddNum(ids...)

	// Cap our channel to 20 messages. This can change later.
	messages := make(chan *imap.Message, 20)
	done := make(chan error, 1)
	go func() {
		done <- mailClient.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, "BODY[]", imap.FetchUid}, messages)
	}()

	for m := range messages {
		// Extra check for only monitoring messages
		if m.Envelope.Sender[0].PersonalName == "DigitalOcean" && strings.HasPrefix(m.Envelope.Subject, "DigitalOcean monitoring triggered:") {
			// Get the actual alert
			r := m.GetBody(&imap.BodySectionName{Peek: true})
			if r == nil {
				// Continue to other messages
				continue
			}

			mr, err := mail.CreateReader(r)
			if err != nil {
				// Continue to other messages
				continue
			}

			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				switch p.Header.(type) {
				case *mail.InlineHeader:
					// This is the message's text (can be plain-text or HTML)
					b, _ := ioutil.ReadAll(p.Body)
					txt := string(b)
					alerts = append(alerts, txt)
				case *mail.AttachmentHeader:
					// We're not interested in attachments
				}
			}

		}
	}

	return alerts, nil
}
