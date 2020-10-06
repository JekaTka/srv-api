package password

import (
	"golang.org/x/crypto/bcrypt"
)

type password struct{}

//NewService create a new fake password
func New() Passwrder {
	return &password{}
}

//Generate a new password
func (p *password) Generate(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//Compare compare two passwords
func (p *password) Compare(p1, p2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2))
	if err != nil {
		return err
	}
	return nil
}
