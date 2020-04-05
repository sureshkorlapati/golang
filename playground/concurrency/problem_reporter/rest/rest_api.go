package rest

import (
	"io"
	"log"
	"net/http"
	"time"
)

type RestHandler struct {
}

func handleAzure(azureChan chan string) {
	time.Sleep(500 * time.Millisecond)
	log.Println("sending azure ....")
	azureChan <- "azure success"
	close(azureChan)
}

func handleKafka(kafkaChan chan string) {
	time.Sleep(500 * time.Millisecond)
	log.Println("sending kafka ....")
	kafkaChan <- "kafka success"
	close(kafkaChan)
}
func (handler RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	azureChan := make(chan string)
	kafkaChan := make(chan string)
	go handleAzure(azureChan)
	go handleKafka(kafkaChan)

	var response string

	// for {
	select {
	case msg := <-azureChan:
		log.Println("receiving azure .... : " + response)
		response = response + " " + msg
	case msg := <-kafkaChan:
		log.Println("receiving kafka .... : " + response)
		response = response + " " + msg
	}
	// }
	io.WriteString(w, response)
}
