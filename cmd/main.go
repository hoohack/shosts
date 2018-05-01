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

func checkArgs(command string, args []string) {
	switch command {
	case "append":
		checkIPErr := shosts.CheckIP(args[0])
		if len(args) != 2 || checkIPErr != nil || !shosts.CheckDomain(args[1]) {
			fmt.Printf("Please input the right args: 'append $ip $domain' eg: append 127.0.0.1 www.baidu.com\n")
			os.Exit(1)
		}
		break
	case "del":
		if len(args) != 1 || !shosts.CheckDomain(args[0]) {
			fmt.Printf("Please input the right args: 'append $ip $domain' eg: append 127.0.0.1 www.baidu.com\n")
			os.Exit(1)
		}
		break
	case "enableGrp":
		if len(args) != 1 {
			fmt.Printf("Please input the right args: 'enableGrp $groupName' eg: enableGrp localhost\n")
			os.Exit(1)
		}
	case "disableGrp":
		if len(args) != 1 {
			fmt.Printf("Please input the right args: 'disableGrp $groupName' eg: disableGrp localhost\n")
			os.Exit(1)
		}
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
	case "enableGrp":
		grpName := args[0]
		hostfile.EnableGroup(grpName)
		break
	case "disableGrp":
		grpName := args[0]
		hostfile.DisableGroup(grpName)
		break
	case "listGrp":
		hostfile.ListCurrentHostsGroup()
		break
	default:
		fmt.Println("Please enter the right command[append|del|list|listGrp|enableGrp]")
	}
}
