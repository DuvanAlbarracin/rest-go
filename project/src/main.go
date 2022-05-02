package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/src/mirrors"
	"time"
)

type response struct {
	FastestURL string
	Latency    time.Duration
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()
	}
	return response{<-urlChan, <-latencyChan}
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter,
		r *http.Request) {
		response := findFastest(mirrors.MirrorList)
        respjson, _ := json.Marshal(response)
        fmt.Printf("%+v", response)
        fmt.Println(respjson)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respjson)
	})
	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
