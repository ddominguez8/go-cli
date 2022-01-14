package main

import (
	"bufio"
	"os"
	"os/exec"
	"fmt"
	"strings"
	"errors"
)

func main () { 
	// Initialize a New Reader to take input.
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		// Reading keyboard input. 
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// Handle execution of input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove newline character.
	input = strings.TrimSuffix(input, "\n")

	// Let's split up the commands from the arguments. 
	args := strings.Split(input, " ")

	// Make a switch & case statements to read change directory commands (Probably refine more?)
	switch args[0] {
	case "cd":

		// You must provide a path. Home is \
		if len(args) < 2 {
			return errors.New("Please provide a path peasant.")
		}
		return os.Chdir(args[1])
	case "exit": 
		os.Exit(0)	
	}

	// Prepare the command to execute. 
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device. 
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}