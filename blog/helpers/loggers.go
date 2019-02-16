package helpers

import "log"

func LogIfError(err error, format string) {
	if err != nil { log.Fatalf(format, err) }
}

