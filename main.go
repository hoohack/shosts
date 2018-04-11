package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
* host配置文件
 */
type Hostfile struct {
	Path  string
	Hosts map[string]Hostname
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
	return &Hostfile{path, make(map[string]Hostname)}
}

func NewHostName(domain string, ip string, enabled bool) *Hostname {
	return &Hostname{domain, ip, enabled}
}

func (h *Hostname) toString() []byte {
	return []byte(h.Domain + " " + h.IP)
}

/*
* 增加一个host记录
 */
func (h *Hostfile) Add(host Hostname) {
	if h.Hosts == nil {
		h.Hosts = make(map[string]Hostname)
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

func appendHost(domain string, ip string) {
	if domain == "" || ip == "" {
		return
	}

	fmt.Println("append" + " " + ip)
	hostname := NewHostName(domain, ip, true)
	err := ioutil.WriteFile(getHostPath(), hostname.toString(), 0644)
	if err != nil {
		fmt.Println(err)
	}
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
	default:
		fmt.Println("Please enter the right command[append]")
	}
}
