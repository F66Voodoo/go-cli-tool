package service

import "fmt"

func PrintStartMessage(name string) {
	if name == "" {
		fmt.Println("go-cli-tool started")
		return
	}
	fmt.Println("go-cli-tool started,", name)
}
