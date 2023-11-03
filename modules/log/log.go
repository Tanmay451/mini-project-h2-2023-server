package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	Logger *log.Logger
)

func init() {
	Logger = log.New()
	Logger.SetOutput(os.Stderr)
	Logger.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyMsg:  "message",
			log.FieldKeyFile: "location",
			log.FieldKeyTime: "timestamp",
		},
	})

}
