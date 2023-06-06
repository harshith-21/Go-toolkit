package main

import "github.com/harshith-21/toolkit"

func main() {
	var tools toolkit.Tools

	_ = tools.CreateDirIfNotExist("./test-dir")
}
