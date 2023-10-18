package sms

import (
	"errors"

	"github.com/twilio/twilio-go"
	tclient "github.com/twilio/twilio-go/client"

	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type Twilio struct {
	client     *twilio.RestClient
	serviceSid string
}

func NewTwilio(cfg Config) *Twilio {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{Username: cfg.Twilio.AccountSid, Password: cfg.Twilio.AuthToken})
	return &Twilio{client: client, serviceSid: cfg.Twilio.ServiceSid}
}

func (t *Twilio) SendVerificationCode(phoneNumber string) (err error) {

	channel := "sms"
	// customFriendly := "CustomFriendlyName" //"https://www.twilio.com/docs/errors/60204"
	// customMessage := "CustomMessage"
	// customCode := "CustomCode"
	_, err = t.client.VerifyV2.CreateVerification(t.serviceSid, &verify.CreateVerificationParams{
		To: &phoneNumber,
		// CustomFriendlyName: &customFriendly,
		// CustomMessage:      &customMessage,
		// CustomCode:         &customCode,
		Channel: &channel,
	})
	if err == nil {
		return nil
	}
	var twilioErr = &tclient.TwilioRestError{}
	if errors.As(err, &twilioErr) && twilioErr.Code == 60220 {
		return ErrNotSupport
	}
	return err
}

func (t *Twilio) VerifyCode(phoneNumber, code string) (err error) {
	resp, err := t.client.VerifyV2.CreateVerificationCheck(t.serviceSid, &verify.CreateVerificationCheckParams{
		To:   &phoneNumber,
		Code: &code,
	})
	if resp != nil && resp.Status != nil && *resp.Status == "approved" {
		return nil
	}
	if err != nil {
		return
	}
	if resp != nil && resp.Status != nil {
		return errors.New("verification failed: " + *resp.Status)
	}
	return errors.New("verification failed")
}
