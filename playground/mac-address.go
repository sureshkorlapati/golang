package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	// func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			fmt.Println(i)
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr := i.HardwareAddr.String()
				fmt.Println(addr)
				break
			}
		}
	}
	return
	// }
}
