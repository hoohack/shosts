package shosts

import (
	"fmt"
	"os"
	"testing"
)

func TestAppend(t *testing.T) {
	filePath := "/home/hoohack/go/src/shosts/sources/hostfile"
	hostfile := NewHostfile(filePath)
	if !hostfile.PathExists(filePath) {
		fmt.Printf("file : %s not exists\n", filePath)
		os.Exit(1)
	}

	hostnameMap := hostfile.ParseHostfile(filePath)
	if len(hostnameMap) != 0 {
		fmt.Printf("file %s is not empty\n", filePath)
		os.Exit(1)
	}

	domain := "localhost"
	ip := "127.0.0.1"

	os.Setenv("GOHOST_FILE", filePath)
	hostfile.AppendHost(domain, ip)

	newHostnameMap := hostfile.ParseHostfile(filePath)
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
