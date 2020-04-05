package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UTC()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	// fmt.Println(t.Format("YYYY-MM-dd HH:mm:ss"))
}
