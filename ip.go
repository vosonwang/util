package util

import (
	"io/ioutil"
	"net"
	"net/http"
)

func GetLocalIP() string {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range ips {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}

func GetExternalIP() (string, error) {
	resp, err := http.Get("http://ifconfig.me")
	if err != nil {
		return "", err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
