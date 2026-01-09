package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/student/go-cli-tool/internal/service"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "go-cli-tool %s\n\n", service.Version())
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "  go-cli-tool [options]")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintln(os.Stderr, "  go run ./cmd/app")
		fmt.Fprintln(os.Stderr, "  go run ./cmd/app --version")
		fmt.Fprintln(os.Stderr, "  go run ./cmd/app --help")
	}

	showVersion := flag.Bool("version", false, "show version and exit")
	name := flag.String("name", "", "print name in start message")
	flag.Parse()

	if *showVersion {
		fmt.Println("go-cli-tool", service.Version())
		return
	}
	service.PrintStartMessage(*name)
	fmt.Println("version", service.Version())
}
