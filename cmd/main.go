package main

import (
	"fmt"
	"github.com/hoohack/shosts"
	"net/url"
	"os"
	"regexp"
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

func checkIP(ip string) bool {
	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	return re.MatchString(ip)
}

func checkDomain(domain string) bool {
	_, errURL := url.Parse(domain)
	return errURL == nil
}

func checkArgs(command string, args []string) {
	switch command {
	case "append":
		if len(args) != 2 || !checkIP(args[0]) || !checkDomain(args[1]) {
			fmt.Printf("Please input the right args: 'append $ip $domain' eg: append 127.0.0.1 www.baidu.com\n")
			os.Exit(1)
		}
		break
	case "del":
		if len(args) != 1 || !checkDomain(args[0]) {
			fmt.Printf("Please input the right args: 'append $ip $domain' eg: append 127.0.0.1 www.baidu.com\n")
			os.Exit(1)
		}
		break
	default:
		break
	}
}

func main() {
	command := getCommand()
	if command == "" {
		fmt.Println("Please enter the right command[append|del|list]")
		os.Exit(1)
	}

	fmt.Println("command: " + command)
	args := getArgs()
	checkArgs(command, args)
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
