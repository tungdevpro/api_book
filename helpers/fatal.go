package helpers

import "log"

func Fatal(err any) {
	if err != nil {
		log.Fatal("err: ", err)
	}
}
