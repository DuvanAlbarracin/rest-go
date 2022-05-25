package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func healthCheck(w http.ResponseWriter, r *http.Request){
    currentTime := time.Now()
    io.WriteString(w, currentTime.String() + "\n")
}

func main(){
    http.HandleFunc("/health", healthCheck)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
