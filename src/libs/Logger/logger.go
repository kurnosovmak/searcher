package Logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

type Logger struct {
	logrus.Logger
}

func New() *Logger {
	var log = Logger{}

	log.Formatter = new(logrus.TextFormatter)                      //default
	log.Formatter.(*logrus.TextFormatter).DisableColors = false    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false // remove timestamp from test output
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	log.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true})

	return &log
}
