package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	if len(os.Args) < 2 {
		fmt.Print(0)
		os.Exit(0)
	}

	str = os.Args[1]

	fmt.Println(str)

	if str != "" {
		s := strings.Fields(str)

		fmt.Print(len(s))
	} else {
		fmt.Print(0)
	}

}
