package global

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

var WebLog *log.Entry
var CoreLog *log.Entry

func createLogger(enableCaller bool) *log.Logger {
	logger := log.New()
	logger.SetReportCaller(enableCaller)
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		ForceQuote:      true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			p, _ := os.Getwd()
			return fmt.Sprintf("%s:%d ", strings.TrimPrefix(f.File, p), f.Line), ""
		},
	})
	return logger
}

func LogInit() {
	WebLog = createLogger(false).WithField("category", "web")
	CoreLog = createLogger(true).WithField("category", "core")
}
