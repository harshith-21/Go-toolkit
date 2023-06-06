package main

import (
	"fmt"
	"github.com/harshith-21/toolkit"
)

func main() {
	var tools toolkit.Tools

	toSlug := "now is the TIME? eh !! 123 &(* 4"

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		fmt.Println("error:: ", err)
	}
	fmt.Println(slugified)
}
