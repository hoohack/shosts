package hosts_test

import (
	"fmt"
	"github.com/hoohack/go-hosts/hosts"
	"os"
)

func TestAppend() {
	filePath := "./hostfile"
	if !hosts.PathExists(filePath) {
		fmt.Printf("file : %s not exists\n", filePath)
		os.Exit(1)
	}

	hostnameMap := hosts.ParseHostFile(filePath)
	if len(hostnameMap) != 0 {
		fmt.Printf("file %s is not empty", filePath)
	}

	domain := "localhost"
	ip := "127.0.0.1"

	os.Setenv()
	AppendHost(domain, ip)

	newHostnameMap := hosts.ParseHostFile(filePath)
	if len(newHostnameMap) == 0 {
		fmt.Println("append test failed")
		os.Exit(1)
	}

	if len(newHostnameMap) == 1 &&
		newHostnameMap[domain] == ip {
		fmt.Println("append test success")
	} else {
		fmt.Println("append test failed")
	}
}
