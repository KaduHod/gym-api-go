package errors

import "log"

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}
