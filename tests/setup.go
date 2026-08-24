package tests

import (
	"errors"
	"os"
	"path"

	FC "github.com/codefin-stack/go-fundconnext"
	godotenv "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	ErrUsernameIsNotDefined = errors.New("username is not defined")
	ErrPasswordIsNotDefined = errors.New("password is not defined")
)

func NewFundConnext() (*FC.FundConnext, error) {
	s, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if err := godotenv.Load(path.Join(s, "../test.env")); err != nil {
		return nil, err
	}
	var username, password string
	var found bool
	if username, found = os.LookupEnv("FC_USERNAME"); !found {
		return nil, ErrUsernameIsNotDefined
	}
	if password, found = os.LookupEnv("FC_PASSWORD"); !found {
		return nil, ErrPasswordIsNotDefined
	}
	logger := log.New()
	logger.SetLevel(log.FatalLevel)
	fc := FC.New(&FC.FCConfiguration{
		Username: username,
		Password: password,
		Env:      "staging",
		Logger:   logger,
	})

	return fc, nil
}
