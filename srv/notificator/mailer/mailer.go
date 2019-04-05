package mailer

import (
	"crypto/tls"
	"os"
	"strconv"
	"time"

	"github.com/micro/go-log"

	"gopkg.in/gomail.v2"
)

// New creates new daemon
func New() chan *gomail.Message {

	send := make(chan *gomail.Message, 50)

	go func() {
		port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
		dialer := gomail.NewDialer(os.Getenv("EMAIL_SERVER"), port, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		var sender gomail.SendCloser
		open := false

		for {
			select {
			case m, ok := <-send:
				if !ok {
					return
				}
				if !open {
					var err error
					if sender, err = dialer.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(sender, m); err != nil {
					log.Logf("Email daemon, error sending email: %v", err)
				}

			case <-time.After(30 * time.Second):
				if open {
					if err := sender.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()

	return send
}
