package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string

	str, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	s := strings.Fields(str)

	fmt.Print(len(s))
}
