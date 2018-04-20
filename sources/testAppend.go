package tests

import (
	"fmt"
	"github.com/hoohack/go-hosts/hosts"
	"os"
)

func TestAppend() {
	filePath := "/home/hoohack/go/go-hosts/tests/hostfile"
	if !hosts.PathExists(filePath) {
		fmt.Printf("file : %s not exists\n", filePath)
		os.Exit(1)
	}

	hostnameMap := hosts.ParseHostFile(filePath)
	if len(hostnameMap) != 0 {
		fmt.Printf("file %s is not empty\n", filePath)
		os.Exit(1)
	}

	domain := "localhost"
	ip := "127.0.0.1"

	os.Setenv("GOHOST_FILE", filePath)
	hosts.AppendHost(domain, ip)

	newHostnameMap := hosts.ParseHostFile(filePath)
	if len(newHostnameMap) == 0 {
		fmt.Println("append test failed")
		os.Exit(1)
	}

	if len(newHostnameMap) == 1 &&
		newHostnameMap[domain].IP == ip {
		fmt.Println("append test success")
	} else {
		fmt.Println("append test failed")
	}
}
