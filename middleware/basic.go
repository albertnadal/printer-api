package middleware

import (
	"time"
	//"log"
	"net/http"
	"strconv"
	"github.com/sirupsen/logrus"
	"github.com/dustin/go-humanize"
	"printer-api/models"
	"printer-api/managers"
)

func InitLogger(config models.Configuration) {
	if(config.Logger.Verbose) {
		customFormatter := new(logrus.TextFormatter)
		customFormatter.TimestampFormat = "2006-01-02 15:04:05"
		logrus.SetFormatter(customFormatter)
		customFormatter.FullTimestamp = true
	}
}

func BasicHandler(f func(http.ResponseWriter, *http.Request, managers.PrinterManager) (int, uint64, error), printerManager managers.PrinterManager, config models.Configuration) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if config.Logger.Verbose {
			// Mode verbose enabled
			beginTime := time.Now()
			statusCode := http.StatusOK
			var responseLength uint64 = 0

			defer func() {
				logger := logrus.WithFields(logrus.Fields{
					"duration":    time.Since(beginTime),
					"status_code": statusCode,
					"remote":      r.RemoteAddr,
					"in_size":     humanize.Bytes(uint64(r.ContentLength)),
					"out_size":    humanize.Bytes(responseLength),
				})

				logger.Info(r.Method + " " + r.URL.RequestURI())
			}()

			statusCode, responseLength, err := f(w, r, printerManager)
			if(config.Logger.Verbose && err != nil) {
				logrus.Error(err)
			}

		} else if !config.Logger.Verbose {
			// Mode verbose disabled
			f(w, r, printerManager)
		}
	})
}

func atouint64(v string) uint64 {
	i64, err := strconv.ParseInt(v, 10, 64)
	if(err != nil) {
		return 0
	} else {
		return uint64(i64)
	}
}
