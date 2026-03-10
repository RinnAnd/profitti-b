package domain

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
