package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleware(originalHandler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Executing middleware before calling originalHandler")
        originalHandler.ServeHTTP(w, r)
        fmt.Println("Executing middleware after calling originalHandler")
    })
}

func handle(w http.ResponseWriter, r *http.Request){
    fmt.Println("Executing mainHandler...")
    w.Write([]byte("OK\n"))
}

func main(){
    originalHandler := http.HandlerFunc(handle)
    http.Handle("/", middleware(originalHandler))
    log.Fatal(http.ListenAndServe(":8000", nil))
}
