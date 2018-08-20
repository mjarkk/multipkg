package gui

import (
	"fmt"
	"log"
	"os"
)

// CheckErr checks if there is an error if so it logs it error
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// FriendlyErr returns a error message without making a user super scared
func FriendlyErr(errors ...interface{}) {
	fmt.Println(errors...)
	os.Exit(0)
}
