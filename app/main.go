package main

import (
	"fmt"
	"github.com/harshith-21/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(5)
	fmt.Println(s)
}
