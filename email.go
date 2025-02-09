package main

import (
	"bytes"
	"net/mail"
	"net/url"
	"strings"
	"text/template"

	"github.com/antonybholmes/go-edb-server-mailer/consts"
	"github.com/antonybholmes/go-mailer"

	"github.com/antonybholmes/go-mailer/mailserver"
)

const JWT_PARAM = "jwt"

//const REDIRECT_URL_PARAM = "redirectUrl"

type EmailBody struct {
	Name       string
	From       string
	Time       string
	Link       string
	DoNotReply string
}

func SendPasswordlessSigninEmail(qe *mailer.RedisQueueEmail) error {

	var file string

	if qe.LinkUrl != "" {
		file = "templates/email/passwordless/web.html"
	} else {
		file = "templates/email/passwordless/api.html"
	}

	go SendEmailWithToken("Passwordless Sign In",
		qe,
		file)

	return nil
}

func SendVerifyEmail(qe *mailer.RedisQueueEmail) error {

	var file string

	if qe.LinkUrl != "" {
		file = "templates/email/verify/web.html"
	} else {
		file = "templates/email/verify/api.html"
	}

	go SendEmailWithToken("Email Address Verification",
		qe,
		file)

	return nil
}

func SendVerifiedEmail(qe *mailer.RedisQueueEmail) error {

	file := "templates/email/verify/verified.html"

	go SendEmailWithToken("Email Address Verified",
		qe,
		file)

	return nil
}

func SendPasswordResetEmail(qe *mailer.RedisQueueEmail) error {

	var file string

	if qe.LinkUrl != "" {
		file = "templates/email/password/reset/web.html"
	} else {
		file = "templates/email/password/reset/api.html"
	}

	go SendEmailWithToken("Password Reset",
		qe,
		file)

	return nil
}

func SendPasswordUpdatedEmail(qe *mailer.RedisQueueEmail) error {

	var file string

	if qe.LinkUrl != "" {
		file = "templates/email/password/switch-to-passwordless.html"
	} else {
		file = "templates/email/password/updated.html"
	}

	go SendEmailWithToken("Password Updated",
		qe,
		file)

	return nil
}

func SendEmailResetEmail(qe *mailer.RedisQueueEmail) error {

	var file string

	if qe.LinkUrl != "" {
		file = "templates/email/email/reset/web.html"
	} else {
		file = "templates/email/email/reset/api.html"
	}

	go SendEmailWithToken("Email Reset",
		qe,
		file)

	return nil
}

func SendEmailUpdatedEmail(qe *mailer.RedisQueueEmail) error {

	file := "templates/email/email/updated.html"

	go SendEmailWithToken("Email Updated",
		qe,
		file)

	return nil
}

func SendAccountCreatedEmail(qe *mailer.RedisQueueEmail) error {

	file := "templates/email/account/created.html"

	go SendEmailWithToken("Account Created",
		qe,
		file)

	return nil
}

func SendAccountUpdatedEmail(qe *mailer.RedisQueueEmail) error {

	file := "templates/email/account/updated.html"

	go SendEmailWithToken("Account Updated",
		qe,
		file)

	return nil
}

func SendEmailWithToken(subject string,
	qe *mailer.RedisQueueEmail,
	file string) error {

	address, err := mail.ParseAddress(qe.To)

	if err != nil {
		return err
	}

	var body bytes.Buffer

	t, err := template.ParseFiles(file)

	if err != nil {
		return err
	}

	var firstName string = ""

	if qe.Name != "" {
		firstName = qe.Name
	} else {
		firstName = strings.Split(address.Address, "@")[0]
	}

	firstName = strings.Split(firstName, " ")[0]

	if qe.LinkUrl != "" {
		linkUrl, err := url.Parse(qe.LinkUrl)

		if err != nil {
			return err
		}

		params, err := url.ParseQuery(linkUrl.RawQuery)

		if err != nil {
			return err
		}

		// after the callback url does some validation, we can
		// goto a different url to make the user experience
		// better
		// this feature is mostly unused since the visit url
		// is normally encoded in the attached jwt to prevent
		// tampering
		//if qe.RedirectUrl != "" {
		//	params.Set(REDIRECT_URL_PARAM, qe.RedirectUrl)
		//}

		if qe.Token != "" {
			params.Set(JWT_PARAM, qe.Token)
		}

		// once we've added extra params, update the
		// raw query again
		linkUrl.RawQuery = params.Encode()

		// the complete url with params
		link := linkUrl.String()

		err = t.Execute(&body, EmailBody{
			Name:       firstName,
			Link:       link,
			From:       consts.NAME,
			Time:       qe.Ttl,
			DoNotReply: consts.DO_NOT_REPLY,
		})

		if err != nil {
			return err
		}
	} else {
		err = t.Execute(&body, EmailBody{
			Name:       firstName,
			Link:       qe.Token,
			From:       consts.NAME,
			Time:       qe.Ttl,
			DoNotReply: consts.DO_NOT_REPLY,
		})

		if err != nil {
			return err
		}
	}

	//log.Debug().Msgf("awhat %v", body.String())

	err = mailserver.SendHtmlEmail(address, subject, body.String())

	if err != nil {
		return err
	}

	return nil
}
