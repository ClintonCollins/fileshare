package sharelogger

import (
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

func callerLog() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
}

func GetLogger(devMode bool) zerolog.Logger {
	callerLog()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if devMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Caller().Logger()
	}
	return zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}
