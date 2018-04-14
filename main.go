package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
* host配置文件
 */
type Hostfile struct {
	Path  string
	Hosts map[string]*Hostname
}

/*
* 单个host属性
 */
type Hostname struct {
	Domain  string
	IP      string
	Enabled bool
}

/*
* host分组
 */
type HostGroup struct {
	Name      string
	GroupFile Hostfile
	Enabled   bool
}

/*
* 实例化
 */
func NewHostfile(path string) *Hostfile {
	return &Hostfile{path, make(map[string]*Hostname)}
}

func NewHostname(domain string, ip string, enabled bool) *Hostname {
	return &Hostname{domain, ip, enabled}
}

func (h *Hostname) toString() string {
	return h.IP + " " + h.Domain
}

/*
* 增加一个host记录
 */
func (h *Hostfile) Add(host *Hostname) {
	if h.Hosts == nil {
		h.Hosts = make(map[string]*Hostname)
	}
	h.Hosts[host.Domain] = host
}

func (h *Hostfile) Delete(host string) {
	delete(h.Hosts, host)
}

/*
* 获取host文件路径
 */
func getHostPath() string {
	path := os.Getenv("GOHOST_FILE")
	if path == "" {
		path = "/etc/hosts"
	}

	return path
}

func addGroup(name string) {

}

func deleteGroup(name string) {

}

func renameGroup(name string) {

}

func enableGroup(name string) {

}

func disableGroup(name string) {

}

func getCommand() string {
	return os.Args[1]
}

func getArgs() []string {
	return os.Args[2:]
}

func appendToFile(filePath string, stringToWrite string) {
	fp, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("failed opening file %s : %s:", filePath, err)
		os.Exit(1)
	}
	defer fp.Close()

	stringToWrite = "\n" + stringToWrite
	_, err = fp.WriteString(stringToWrite)
	if err != nil {
		fmt.Println("failed append string: %s:", filePath, err)
		os.Exit(1)

	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil || os.IsNotExist(err) {
		return true
	}

	return false
}

func parseHostFile(path string) map[string]*Hostname {
	if !PathExists(path) {
		fmt.Println("path %s is not exists", path)
		os.Exit(1)
	}

	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file %s fail: %s", path, err)
		os.Exit(1)
	}

	hostnameArr := strings.Split(string(fileContents[:]), "\n")
	hostnameMap := make(map[string]*Hostname)
	for _, val := range hostnameArr {
		if len(val) == 0 || val == "\r\n" {
			continue
		}
		tmpHostnameArr := strings.Split(val, " ")
		tmpHostname := NewHostname(tmpHostnameArr[1], tmpHostnameArr[0], true)
		hostnameMap[tmpHostname.Domain] = tmpHostname
	}

	fmt.Println(hostnameMap)
	return hostnameMap
}

func appendHost(domain string, ip string) {
	if domain == "" || ip == "" {
		return
	}

	fmt.Println("append" + " " + ip)
	hostname := NewHostname(domain, ip, true)
	appendToFile(getHostPath(), hostname.toString())
}

func writeToFile(hostnameMap map[string]*Hostname, path string) {
	if !PathExists(path) {
		fmt.Println("path %s is not exists", path)
		os.Exit(1)
	}

	fp, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer fp.Close()

	for _, mapVal := range hostnameMap {
		fp.WriteString(mapVal.toString())
	}
}

func deleteDomain(domain string) {
	if domain == "" {
		return
	}

	currHostsMap := parseHostFile(getHostPath())

	if len(currHostsMap) == 0 || currHostsMap[domain] == nil {
		return
	}

	delete(currHostsMap, domain)
	writeToFile(currHostsMap, getHostPath())
}

func main() {
	command := getCommand()
	if command == "" {
		os.Exit(1)
	}

	fmt.Println("command: " + command)
	args := getArgs()
	switch command {
	case "append":
		domain := args[0]
		ip := args[1]
		appendHost(domain, ip)
		break
	case "del":
		domain := args[0]
		deleteDomain(domain)
		break
	default:
		fmt.Println("Please enter the right command[append]")
	}
}
