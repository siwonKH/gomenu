package handler

import (
	"fmt"
	"log"
)

func HandleErr(err string) error {
	log.Println(err)
	return fmt.Errorf("%s", err)
}
