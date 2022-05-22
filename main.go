package main

import (
	"os"

	"github.com/cfelegy/aggy/src/backend"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	// TODO: infer logrus level from environment variable
	logrus.SetLevel(logrus.TraceLevel)
}

func main() {
	logrus.Infoln("backend starting...")
	backend.Start()
}
