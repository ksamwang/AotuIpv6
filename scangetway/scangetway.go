package scangetway

import (
	"fmt"
	"net"
)

func Reipv()(string,int)  {
	ipv6:=""
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "错误",-1
	}
	for _, add := range address {
		ipnet, _:= add.(*net.IPNet)
		ipv6 = ipnet.IP.String()
		if ipv6[0:3] == "240"{
			return  ipv6,0
		}
	}
	return  "找不到IPV6",-1
}
