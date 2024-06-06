package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

const (
	reset = "\033[0m"
	green = "\033[32m"
	red   = "\033[31m"
	blue  = "\033[36m"
)

func main() {
	clear()
	motd()

	reader := bufio.NewReader(os.Stdin)

	usr, err := user.Current()
	if err != nil {
		fmt.Println("User error: ", err)
	}

	for {
		cwd, err := os.Getwd()

		if err != nil {
			fmt.Println("Directory error: ", err)
		}
		fmt.Printf(green+"%v"+reset+"@"+blue+"%v "+reset+"> ", usr.Username, cwd)

		input, err := reader.ReadString('\n')
		if err != nil {
			printError(err)
		}

		input = strings.TrimSpace(input)

		err = execInput(input)
		if err != nil {
			printError(err)

		}
	}

}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "help":
		fmt.Printf("Welcome to help!\nThis shell was build by izac, and it's more than likely that you're izac.\nIf you don't remember how to use it, look at the code!\n")
		return nil
	case "exit":
		fmt.Println("Goodbye!")
		os.Exit(0)

	case "info":
		fmt.Printf("Build by Izac Peterson\nizacpeterson.com\ngithub.com/izacpeterson\n")
		return nil

	case "hello":
		usr, err := user.Current()
		if err != nil {
			fmt.Println("User error: ", err)
		}
		fmt.Printf("Hello, %v.\n", usr.Username)
		return nil

	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func printError(err error) {
	fmt.Println(red+"error: ", err)
}

func motd() {
	motdFile := "motd.txt"
	content, err := os.ReadFile(motdFile)
	if err != nil {
		printError(err)
	}
	fmt.Println(string(content))

}
func clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
