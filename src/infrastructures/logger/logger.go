package loggers

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// const (
// 	PROD = "production"
// )

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
