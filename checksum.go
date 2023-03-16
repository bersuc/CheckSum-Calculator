/*
go run checksum.go "8=FIX.4.2|9=49|35=5|34=1|49=ARCA|52=20150916-04:14:05.306|56=TW|"
Reference : https://stackoverflow.com/questions/32708068/how-to-calculate-checksum-in-fix-manually
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	var arg1 = os.Args[1]
	var soma = 0
	for _, c := range arg1 {
		if string(c) == "|" {
			soma += 1
			continue
		} else {
			soma += int(c)
		}
	}
	fmt.Printf("%s10=%d", arg1, soma%256)

}
