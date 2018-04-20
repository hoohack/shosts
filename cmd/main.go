package main

import (
	"fmt"
	"github.com/hoohack/shosts"
	"os"
)

func getCommand() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		return ""
	}
}

func getArgs() []string {
	return os.Args[2:]
}

func main() {
	command := getCommand()
	if command == "" {
		fmt.Println("Please enter the right command[append|del|list]")
		os.Exit(1)
	}

	fmt.Println("command: " + command)
	args := getArgs()
	filePath := "/etc/hosts"
	hostfile := shosts.NewHostfile(filePath)
	switch command {
	case "append":
		domain := args[1]
		ip := args[0]
		hostfile.AppendHost(domain, ip)
		break
	case "del":
		domain := args[0]
		hostfile.DeleteDomain(domain)
		break
	case "list":
		hostfile.ListCurrentHosts()
		break
	default:
		fmt.Println("Please enter the right command[append|del|list]")
	}
}
