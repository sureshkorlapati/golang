package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	currentTime := now.Unix()
	log.Println(fmt.Sprintf("%v", currentTime))
}
