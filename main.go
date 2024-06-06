package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input Error: ", err)
		}

		err = execInput(input)
		if err != nil {
			fmt.Println("Input Error: ", err)

		}
	}

}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)

	if args[0] == "cd" {
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
