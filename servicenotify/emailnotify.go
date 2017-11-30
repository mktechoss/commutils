package servicenotify

import (
	"github.com/pkg/errors"
	"net/smtp"
	"strings"
	"sync"
	"time"
)

type EmailNotify struct {
	username     string
	password     string
	host         string
	to           string
	lastsendtime int64
	timelock     sync.Mutex
}

func NewEmailInstance(username, password, host, to string) Notify {
	notify := &EmailNotify{
		username: username,
		password: password,
		host:     host,
		to:       to,
	}
	return notify
}

func (this *EmailNotify) Send(title, content string) error {
	this.timelock.Lock()
	defer this.timelock.Unlock()
	if time.Now().Unix()-this.lastsendtime < 10 {
		return errors.New("too fast.")
	}
	this.lastsendtime = time.Now().Unix()

	hp := strings.Split(this.host, ":")
	auth := smtp.PlainAuth("", this.username, this.password, hp[0])
	content_type := "Content-Type: text/plain" + "; charset=UTF-8"
	msg := []byte("To: " + this.to + "\r\nFrom: " +
		this.username + "<" + this.username + ">\r\nSubject: " +
		title + "\r\n" + content_type + "\r\n\r\n" + content)
	send_to := strings.Split(this.to, ";")
	err := smtp.SendMail(this.host, auth, this.username, send_to, msg)
	return err
}
