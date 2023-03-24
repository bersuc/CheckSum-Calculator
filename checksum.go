/*
Add one message per line as example on messages.txt.
go run checksum.go
Reference : https://stackoverflow.com/questions/32708068/how-to-calculate-checksum-in-fix-manually
*/

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	readFile()
}

// Read the messages.txt file and send to Chacksum calculation function
func readFile() {
	mydir, err := os.Getwd()
	check(err)
	infos, err := os.ReadFile(mydir + "\\messages.txt")
	check(err)
	if len(infos) != 0 {
		messages := strings.Split(string(infos), "\r\n")
		calcCheckSum(messages)
	} else {
		err := errors.New("No line in file. Please follow the steps on docs")
		check(err)
	}

}

// Get the slice of strings from the messages.txt file and print on console
func calcCheckSum(args []string) {
	for _, message := range args {
		var soma = 0
		for _, c := range message {
			if string(c) == "|" {
				soma += 1
				continue
			} else {
				soma += int(c)
			}
		}
		fmt.Printf("%s10=%d\n", message, soma%256)
	}
}

// Send an error to this function.
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
