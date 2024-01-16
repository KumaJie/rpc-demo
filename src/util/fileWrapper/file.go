package fileWrapper

import (
	"fmt"
	"net"
)

var fsIP string

func init() {
	fsIP = GetFSPath()
}

func GetFSPath() string {
	// 获取所有网络接口
	interfaces, _ := net.Interfaces()
	var ip string
	// 遍历所有网络接口
	for _, iface := range interfaces {
		if iface.Name == "WLAN" {
			addrs, _ := iface.Addrs()
			if addr, ok := addrs[len(addrs)-1].(*net.IPNet); ok {
				ip = addr.IP.String()
			}
		}
	}
	return fmt.Sprintf("http://%s:8080", ip)

}

func GetFullPath(filename string) string {
	return fmt.Sprintf("%s/feed/%s", fsIP, filename)
}
