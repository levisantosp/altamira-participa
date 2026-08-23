package utils

import "log"

func LogErr(err error) error {
	log.Println(err)
	return err
}
