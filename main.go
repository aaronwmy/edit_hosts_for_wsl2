package main

import (
	"os"
	"bufio"
	"net"
	"strings"
)

var hostsPath = "/mnt/c/Windows/System32/drivers/etc/hosts"

func main() {
	ip:=GetOutboundIP()
	newContent:=GetNewContent(ip)
	WriteNewContent(newContent)
	
}

// 获得服务器ip
func GetOutboundIP() string {
    conn, _ := net.Dial("udp", "8.8.8.8:80")
    defer conn.Close()
 
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP.String()
}

// 获得新的hosts文件内容
func GetNewContent(ip string) string {
	var newContent strings.Builder
	file, _ := os.Open(hostsPath)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		linStr := fileScanner.Text()
		if len(linStr)>0 {
			if linStr[0:1]!="#" {
				linArr := strings.Split(linStr," ")
				if len(linArr)==2 {
					linStr = ip + " " + linArr[1]
				}
			}
		}
		newContent.WriteString(linStr + "\r\n")
	}
	return newContent.String()
}

// 覆盖写入hosts
func WriteNewContent(content string) {
	f, _ := os.OpenFile(hostsPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	n, _ := f.Seek(0, os.SEEK_END)
    f.WriteAt([]byte(content), n)
    defer f.Close()
}