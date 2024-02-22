package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define command-line flags
	option := flag.String("option", "", "Choose one of the following options: connect, publish, subscribe, qos, retained, lwt, tls")
	flag.Parse()

	// Validate the provided option
	switch *option {
	case "connect":
		runGoFile("connect.go")
	case "publish":
		runGoFile("publish.go")
	case "subscribe":
		runGoFile("subscribe.go")
	case "qos":
		runGoFile("qos.go")
	case "retained":
		runGoFile("retained.go")
	case "lwt":
		runGoFile("lwt.go")
	case "tls":
		runGoFile("tls.go")
	default:
		fmt.Println("Invalid option. Please choose one of the following options: connect, publish, subscribe, qos, retained, lwt, tls")
		os.Exit(1)
	}
}

func runGoFile(fileName string) {
	cmd := exec.Command("go", "run", "examples/"+fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running %s: %v\n", fileName, err)
		os.Exit(1)
	}
}
