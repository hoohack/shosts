package shosts

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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
	Comment string // 注释
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

func NewHostname(comment string, domain string, ip string, enabled bool) *Hostname {
	return &Hostname{comment, domain, ip, enabled}
}

func (h *Hostname) toString() string {
	if len(h.Comment) > 0 {
		h.Comment += "\n"
	}
	return h.Comment + h.IP + " " + h.Domain + "\n"
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

func (h *Hostfile) AddGroup(grpName string) {
	curWd, _ := os.Getwd()
	grpFilePath := curWd + "/sources/group/" + grpName
	if !h.PathExists(grpFilePath) {
		fmt.Printf("group %s's file not exists, please add group file first\n", grpFilePath)
	} else {
		fmt.Println("yes")
	}

	//if !h.grpFileValid(grpName) {
	//fmt.Printf("group file %s's syntax is wrong, please reedit.[ip] [domain] per line\n", grpFilePath)
	//}

	//hostnameMap := h.ParseHostfile(grpFilePath)
	//for domain, ip := range hostnameMap {
	//h.AppendHost(domain, ip)
	/*}*/
}

func deleteGroup(name string) {

}

func renameGroup(name string) {

}

func enableGroup(name string) {

}

func disableGroup(name string) {

}

func appendToFile(filePath string, stringToWrite string) {
	fp, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("failed opening file %s : %s\n", filePath, err)
		os.Exit(1)
	}
	defer fp.Close()

	stringToWrite = "\n" + stringToWrite + "\n"
	_, err = fp.WriteString(stringToWrite)
	if err != nil {
		fmt.Printf("failed append string: %s: %s\n", filePath, err)
		os.Exit(1)
	}
}

func (h *Hostfile) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		return true
	}

	return false
}

func IsEmptyLine(str string) bool {
	re := regexp.MustCompile(`^\s*$`)

	return re.MatchString(str)
}

func (h *Hostfile) ParseHostfile(path string) map[string]*Hostname {
	if !h.PathExists(path) {
		fmt.Printf("path %s is not exists", path)
		os.Exit(1)
	}

	fp, fpErr := os.Open(path)
	if fpErr != nil {
		fmt.Printf("open file '%s' failed\n", path)
		os.Exit(1)
	}
	defer fp.Close()

	br := bufio.NewReader(fp)
	hostnameMap := make(map[string]*Hostname)
	curComment := ""
	for {
		str, rErr := br.ReadString('\n')
		if rErr == io.EOF {
			break
		}
		if 0 == len(str) || str == "\r\n" || IsEmptyLine(str) {
			continue
		}

		if str[0] == '#' {
			// 处理注释
			curComment += str
			continue
		}
		tmpHostnameArr := strings.Split(str, " ")
		domain := strings.Trim(tmpHostnameArr[1], " \n")
		tmpHostname := NewHostname(curComment, domain, tmpHostnameArr[0], true)
		hostnameMap[tmpHostname.Domain] = tmpHostname

		curComment = ""
	}

	return hostnameMap
}

func (h *Hostfile) AppendHost(domain string, ip string) {
	if domain == "" || ip == "" {
		return
	}

	hostname := NewHostname("", domain, ip, true)
	appendToFile(getHostPath(), hostname.toString())
}

func (h *Hostfile) writeToFile(hostnameMap map[string]*Hostname, path string) {
	if !h.PathExists(path) {
		fmt.Printf("path %s is not exists", path)
		os.Exit(1)
	}

	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fp.Close()

	for _, mapVal := range hostnameMap {
		_, writeErr := fp.WriteString(mapVal.toString())
		if writeErr != nil {
			fmt.Println(writeErr)
			os.Exit(1)
		}
	}
}

func (h *Hostfile) DeleteDomain(domain string) {
	if domain == "" {
		return
	}

	currHostsMap := h.ParseHostfile(getHostPath())

	if len(currHostsMap) == 0 || currHostsMap[domain] == nil {
		fmt.Printf("domain %s not exist\n", domain)
		return
	}

	delete(currHostsMap, domain)
	h.writeToFile(currHostsMap, getHostPath())
}

func (h *Hostfile) ListCurrentHosts() {
	currHostsMap := h.ParseHostfile(getHostPath())
	if len(currHostsMap) == 0 {
		return
	}

	for _, mapVal := range currHostsMap {
		fmt.Println(mapVal.toString())
	}
}
