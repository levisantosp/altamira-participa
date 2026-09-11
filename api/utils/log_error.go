package utils

import "log"

func LogErr(genericErr error, err error) error {
	log.Println(err)
	return genericErr
}
