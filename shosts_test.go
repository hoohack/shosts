package shosts

import (
	"fmt"
	"os"
	"testing"
)

func TestAppend(t *testing.T) {
	filePath := "/Users/hoohack/Projects/Go/src/shosts/sources/hostfile"
	var _, err = os.Stat(filePath)
	if os.IsExist(err) {
		fmt.Printf("file %s already exists, remove first\n", filePath)
		delErr := os.Remove(filePath)
		if delErr != nil {
			fmt.Printf("file %s delete failed\n", filePath)
			os.Exit(1)
		}

		fmt.Printf("==> done remove file: %s\n", filePath)
	}

	var file, createErr = os.Create(filePath)
	if createErr != nil {
		fmt.Println(createErr)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Printf("==> done creating file: %s\n", filePath)

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

	fmt.Printf("test finish. start deleting file: %s \n", filePath)
	delErr := os.Remove(filePath)
	if delErr != nil {
		fmt.Printf("file %s delete failed\n", filePath)
		os.Exit(1)
	}
	fmt.Printf("==> done deleted file: %s\n", filePath)

	os.Setenv("GOHOST_FILE", "/etc/hosts")
}

func TestDelete(t *testing.T) {
	filePath := "/home/hoohack/go/src/shosts/sources/testdelete"
	var _, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		var file, err = os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
	} else {
		fmt.Printf("file %s already exists, please remove first\n", filePath)
	}

	fmt.Printf("==> done creating file: %s\n", filePath)
	hostfile := NewHostfile(filePath)
	if !hostfile.PathExists(filePath) {
		fmt.Printf("file : %s not exists please create first\n", filePath)
		os.Exit(1)
	}

	hostnameMap := hostfile.ParseHostfile(filePath)
	if len(hostnameMap) != 0 {
		fmt.Printf("file %s is not empty\n", filePath)
		os.Exit(1)
	}

	domain := "test.delete.com"
	ip := "127.0.1.2"

	os.Setenv("GOHOST_FILE", filePath)
	hostfile.AppendHost(domain, ip)

	newHostnameMap := hostfile.ParseHostfile(filePath)
	if len(newHostnameMap) == 0 {
		fmt.Println("append failed")
		os.Exit(1)
	}

	hostfile.DeleteDomain("test.delete.com")
	hostnameMapAfterDelete := hostfile.ParseHostfile(filePath)
	if len(hostnameMapAfterDelete) == 0 {
		fmt.Println("delete test success")
	} else {
		fmt.Println("delete test failed")
	}

	delErr := os.Remove(filePath)
	if delErr != nil {
		fmt.Printf("file %s delete failed\n", filePath)
	}

	fmt.Println("==> done deleting file")

	os.Setenv("GOHOST_FILE", "/etc/hosts")
}
