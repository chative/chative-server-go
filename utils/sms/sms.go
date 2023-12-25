package sms

import (
	"math/rand"
	"strings"
	"time"
)

type Config struct {
	Twilio struct {
		AccountSid string
		AuthToken  string
		ServiceSid string
	}
	B2M struct {
		AppID  string
		AppKey string
	}
}

type SMS struct {
	Twilio *Twilio
	// B2M    *B2M
}

var (
	sms *SMS
)

func Init(cfg Config) *SMS {
	sms = &SMS{
		Twilio: NewTwilio(cfg),
		// B2M:    NewB2M(cfg),
	}
	rand.Seed(time.Now().UnixMilli())
	return sms
}

func (s *SMS) SendCode(phoneNumber string) error {
	if strings.HasPrefix(phoneNumber, "+86") {
		return ErrNotSupport
		// return s.B2M.SendVerificationCode(phoneNumber)
	}
	return s.Twilio.SendVerificationCode(phoneNumber)
}

func (s *SMS) VerifyCode(phoneNumber, code string) error {
	if strings.HasPrefix(phoneNumber, "+86") {
		return ErrNotSupport
		// return s.B2M.VerifyCode(phoneNumber, code)
	}
	return s.Twilio.VerifyCode(phoneNumber, code)
}

func SendCode(phoneNumber string) error {
	return sms.SendCode(phoneNumber)
}
func VerifyCode(phoneNumber, code string) error {
	return sms.VerifyCode(phoneNumber, code)
}
