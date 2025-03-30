package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type commandFunc func(args []string)

func main() {
	commands := make(map[string]commandFunc)

	commands["exit"] = func(args []string) {
		if len(args) > 0 && args[0] == "0" {
			os.Exit(0)
		}
		os.Exit(0)
	}

	commands["echo"] = func(args []string) {
		if len(args) > 0 {
			fmt.Println(strings.Join(args, " "))
		}
	}

	commands["type"] = func(args []string) {
		if len(args) > 0 {
			if _, ok := commands[args[0]]; ok {
				fmt.Printf("%s is a shell builtin\n", args[0])
			} else {
				fmt.Printf("%s: not found\n", args[0])
			}
		}
	}

	commands["ls"] = func(args []string) {
		if len(args) > 0 {
			for dir := range args {
				entries, _ := os.ReadDir(args[dir])
				for _, e := range entries {
					fmt.Println(e.Name())
				}
			}

		} else if len(args) == 0 {
			entries, _ := os.ReadDir("./")
			for _, e := range entries {
				fmt.Println(e.Name())
			}
		}
	}

	commands["cat"] = func(args []string) {
		if len(args) == 0 {
			file, err := os.ReadFile(args[0])
			if err != nil {
				fmt.Printf("cat: %s: No such file or directory\n", args)
			}
			fmt.Println(string(file))
		} else if len(args) > 1 {
			for path := range args {
				file, err := os.ReadFile(args[path])
				if err != nil {
					fmt.Printf("cat: %s: No such file or directory\n", args)
				}
				fmt.Println(string(file))
			}
		}

	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}

		parts := strings.Fields(command)
		commandName := parts[0]
		args := parts[1:]

		if cmd, exists := commands[commandName]; exists {
			cmd(args)
		} else {
			fmt.Printf("%s: command not found\n", commandName)
		}
	}
}
