package main

import (
	"flag"
	"fmt"

	"github.com/student/go-cli-tool/internal/service"
)

func main() {
	showVersion := flag.Bool("version", false, "show version")
	flag.Parse()

	if *showVersion {
		fmt.Println("go-cli-tool", service.Version())
		return
	}
	service.PrintStartMessage()
	fmt.Println("version", service.Version())
}
