package smtp_client

import (
	"errors"
	"log"
	"net/textproto"
	"time"

	"github.com/go_email_client/pkg/types"
	"github.com/jordan-wright/email"
)

func (sc *SmtpClients) SendMail(
	to []string,
	subject string,
	htmlContent string,
	overrides *types.HeaderOverrides,
) error {
	sc.counter += 1
	if len(sc.connectionPool) < 1 {
		sc.connectionPool = initConnectionPool(sc.servers)
		if len(sc.connectionPool) < 1 {
			return errors.New("no servers defined")
		}
	}

	index := sc.counter % len(sc.connectionPool)
	selectedServer := sc.connectionPool[index]

	From := sc.servers.From
	Sender := sc.servers.Sender
	ReplyTo := sc.servers.ReplyTo

	if overrides != nil {
		if overrides.From != "" {
			From = overrides.From
		}
		if overrides.Sender != "" {
			Sender = overrides.Sender
		}

		if overrides.NoReplyTo {
			ReplyTo = []string{}
		} else if len(overrides.ReplyTo) > 0 {
			ReplyTo = overrides.ReplyTo
		}
	}

	e := &email.Email{
		To:      to,
		From:    From,
		Sender:  Sender,
		ReplyTo: ReplyTo,
		Subject: subject,
		HTML:    []byte(htmlContent),
		Headers: textproto.MIMEHeader{},
	}
	err := selectedServer.Send(e, time.Second*time.Duration(sc.servers.Servers[index].SendTimeout))

	if err != nil {
		// close and try to reconnect
		log.Printf("error when trying to send email: %v", err)
		pool, errReconnect := connectToPool(sc.servers.Servers[index])
		if errReconnect != nil {
			log.Printf("cannot reconnect pool for %v", sc.servers.Servers[index])
		} else {
			log.Printf("successfully reconnected to %v", sc.servers.Servers[index])
			sc.connectionPool[index] = *pool
		}
	}
	return err
}
