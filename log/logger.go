package log

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var LoggingLevel string
var Verbose bool
var LogWithTs bool

func Logger() log.Logger {
	logger := log.NewLogfmtLogger(os.Stdout)

	if LogWithTs {
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	}

	if Verbose {
		logger = level.NewFilter(logger, level.AllowAll())
	} else {
		switch LoggingLevel {
		default:
			logger = level.NewFilter(logger, level.AllowInfo())
		case "debug":
			logger = level.NewFilter(logger, level.AllowDebug())
			logger = log.With(logger, "caller", log.DefaultCaller)
		case "info":
			logger = level.NewFilter(logger, level.AllowInfo())
		case "warn":
			logger = level.NewFilter(logger, level.AllowWarn())
		case "error":
			logger = level.NewFilter(logger, level.AllowError())
		}
	}
	return logger
}
