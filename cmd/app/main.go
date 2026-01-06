package main

import (
	"fmt"

	"github.com/student/go-cli-tool/internal/service"
)

func main() {
	service.PrintStartMessage()
	fmt.Println("version", service.Version())
}
