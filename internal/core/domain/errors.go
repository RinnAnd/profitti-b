package domain

import "errors"

type CredentialsError struct {
	msg string
}

func (c *CredentialsError) Error() string {
	return c.msg
}

func CrdntlsErr(msg string) *CredentialsError {
	return &CredentialsError{
		msg: msg,
	}
}

var User404 = errors.New("user not found")
var Token409 = errors.New("no access token provided")
