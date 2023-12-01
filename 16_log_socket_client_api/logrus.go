package main

import (
	log "github.com/sirupsen/logrus"
)

func logrusMethod() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	// to only show `Error` level
	// log.SetLevel(log.ErrorLevel)

	log.Info("Hello, you!")

	log.WithFields(log.Fields{
		"Name": "Wolf",
		"Diet": "Lamb meat",
	}).Info("Animal Fact")

	log.WithFields(log.Fields{
		"Name":     "Evil spirit",
		"Distance": "30 meter",
	}).Warn("Careful!")

	log.WithFields(log.Fields{
		"Culprit": "Evil spirit",
		"Assault": "Haunted",
	}).Error("You are attacked!")

	// to pass log with `context`
	contextLogger := log.WithFields(log.Fields{
		"context": "parent log",
	})

	contextLogger.Info("Info child")
	contextLogger.Error("Error child")
}
