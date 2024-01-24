package fileWrapper

import (
	"fmt"
	"rpc-douyin/src/config"
)

var fsIP string

func init() {
	fsIP = fmt.Sprintf("http://%s:%d", config.Cfg.Server.Host, config.Cfg.Server.Port)
}
func GetFullPath(filename string) string {
	return fmt.Sprintf("%s/feed/%s", fsIP, filename)
}
