package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "naufal").Info("Hello World")
	logger.WithField("username", "hakim").
		WithField("name", "Naufal Hakim").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "naufal",
		"name":     "Naufal Hakim",
	}).Info("Hello World")
}
