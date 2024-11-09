package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/manendrapalsingh/gorunner/utils"
)

var serverCmd *exec.Cmd

func main() {

	//check that there two argument or not
	if len(os.Args) < 2 {

		fmt.Println("Please use in this format : gorunner <main.go>")
		os.Exit(1)

	}

	// to get the path of the file
	path, err := utils.GetPath()
	if err != nil {

		fmt.Println("Failed to get the path ")
	}

	// start the server
	startServer(os.Args[1])

	// for checking for the file change
	for {

		time.Sleep(500 * time.Millisecond)
		changed, err := utils.CheckForChanges(path)
		if err != nil {
			fmt.Println("error while checking the file ", err)
			continue
		}

		if changed {
			restartServer(os.Args[1])
		}
	}

}

func startServer(fileName string) {
	fmt.Println("Starting a go runner for file:", fileName)
	serverCmd = exec.Command("go", "run", fileName)
	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = os.Stderr

	err := serverCmd.Start()
	if err != nil {
		fmt.Println("Error while running the go file:", fileName, err)
	}
}

func restartServer(filename string) {

	fmt.Println("Re running the go file :", filename)

	if serverCmd != nil && serverCmd.Process != nil {
		serverCmd.Process.Kill()
	}

	startServer(filename)
}
