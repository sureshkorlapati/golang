package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Create("/var/log/frontline.log")
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	for i := 0; i < 100; i++ {
		f.WriteString("{\"lineno\":\"" + strconv.Itoa(i) + "\", \"File\":\"logger.go\",\"Function\":\"logger.Info\",\"Line\":71,\"LogType\":0,\"ServiceName\":\"SIMBA\",\"level\":\"info\",\"msg\":\"Message20-Input Init Data filepath: /config/SimbaData.json\",\"time\":\"2020-01-29T19:17:50.868056798Z\"}\n")
	}
}
