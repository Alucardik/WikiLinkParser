package error_utils

import "log"

func NotifyOnError(msg string, err error) {
	if err != nil {
		if err != nil {
			log.Printf("%s: %s", msg, err)
		}
	}
}

func FailOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
