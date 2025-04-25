package session

import (
	"math/rand"
	"time"
)

type Session struct {
	Code string
}

func generateRandomCode() string {
	rand.Seed(time.Now().Unix())

	const length = 6
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}

func CreateSession() (*Session, error) {
	session := &Session{
		Code: generateRandomCode(),
	}

	return session, nil
}

func (s *Session) GetSessionCode() (string, error) {
	return s.Code, nil
}
