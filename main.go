package main

import (
	"fmt"
	"log"
	"os"
	"pf/account"
	"pf/commands"
)

func main() {
	fmt.Println("Hello world")

	account.LoadAccountList()

	log.SetFlags(0)
	err := commands.Execute(os.Args[1:])
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		// Save only if there were no errors.
		account.SaveAccountList()
	}
}
