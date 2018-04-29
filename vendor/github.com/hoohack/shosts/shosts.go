package shosts

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/url"
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
		fmt.Printf("group %s's file: %s not exists, please add group file first\n", grpName, grpFilePath)
		os.Exit(1)
	}

	grpHostnameMap, parseErr := h.ParseHostfile(grpFilePath)
	if parseErr != nil {
		fmt.Println("parse file failed" + parseErr.Error())
	}

	if len(grpHostnameMap) == 0 {
		fmt.Printf("group file %s is empty,please add host first\n", grpFilePath)
		os.Exit(1)
	}

	for _, val := range grpHostnameMap {
		appendToFile(getHostPath(), val)
	}
}

func deleteGroup(name string) {

}

func renameGroup(name string) {

}

func appendToFile(filePath string, hostname *Hostname) {
	fp, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("failed opening file %s : %s\n", filePath, err)
		os.Exit(1)
	}
	defer fp.Close()

	_, err = fp.WriteString(hostname.toString())
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

func TrimWS(str string) string {
	return strings.Trim(str, " \n\t")
}

func CheckIP(ip string) bool {
	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	return re.MatchString(ip)
}

func CheckDomain(domain string) bool {
	_, errURL := url.Parse(domain)
	return errURL == nil
}

func (h *Hostfile) ParseHostfile(path string) (map[string]*Hostname, error) {
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
		tmpHostnameArr := strings.Fields(str)
		curDomain := TrimWS(tmpHostnameArr[1])
		if !CheckDomain(curDomain) {
			return hostnameMap, errors.New(" file contain error domain" + curDomain)
		}
		curIP := TrimWS(tmpHostnameArr[0])
		if !CheckIP(curIP) {
			return hostnameMap, errors.New(" file contain error ip" + curIP)
		}
		tmpHostname := NewHostname(curComment, curDomain, curIP, true)
		hostnameMap[tmpHostname.Domain] = tmpHostname

		curComment = ""
	}

	return hostnameMap, nil
}

func (h *Hostfile) AppendHost(domain string, ip string) {
	if domain == "" || ip == "" {
		return
	}

	hostname := NewHostname("", domain, ip, true)
	appendToFile(getHostPath(), hostname)
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

	currHostsMap, parseErr := h.ParseHostfile(getHostPath())
	if parseErr != nil {
		fmt.Println("parse file failed" + parseErr.Error())
		return
	}

	if len(currHostsMap) == 0 || currHostsMap[domain] == nil {
		fmt.Printf("domain %s not exist\n", domain)
		return
	}

	delete(currHostsMap, domain)
	h.writeToFile(currHostsMap, getHostPath())
}

func (h *Hostfile) ListCurrentHosts() {
	currHostsMap, parseErr := h.ParseHostfile(getHostPath())
	if parseErr != nil {
		fmt.Println("parse file failed" + parseErr.Error())
		return
	}
	if len(currHostsMap) == 0 {
		return
	}

	for _, mapVal := range currHostsMap {
		fmt.Print(mapVal.toString())
	}
}
