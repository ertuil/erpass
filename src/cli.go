package  main

import (
    "fmt"
    // "os"
	// "strings"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
)

func generateCLI(ac string) {
	fmt.Printf("Master key:")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err != nil {
        fmt.Println("\nCan not readthe password.")
	} else {
		fmt.Println("\nGenerate the password.")
	}
	password := string(bytePassword)
	js := make(map[string]string)

	js["mk"]= password
	js["account"]  = ac
	js["length"]  = "16"
	js["count"]  = "1"
	js["style"]  = "1"
	pa :=  generatePassword(js,"")
	fmt.Println(pa)
}