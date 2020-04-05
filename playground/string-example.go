package main

import (
	"fmt"
	"strings"
)

func main() {
	string := "2020-03-10 19:40:39.483 /usr/bin/simbaclient-go[716]: {\"Function\":\"mrclient.InitMRClient\",\"LogType\":0,\"ServiceName\":\"SIMBA\",\"level\":\"info\",\"msg\":\"InitMRClient HEADLESS fromChannel:11111105 destChannel:4242\",\"time\":\"2020-03-10T19:40:39.482673902Z\"}"

	index := strings.Index(string, "{")
	fmt.Println(string[index:])
}
