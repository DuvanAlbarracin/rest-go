package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type city struct{
    Name string
    Area uint64
}

func postHandler(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST"{
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("405 - Method Not Allowed\n"))
        return
    }

    var currCity city
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&currCity)
    if err != nil{
        panic(err)
    }
    defer r.Body.Close()

    fmt.Printf("City: %s - Area in sq miles: %d", currCity.Name, currCity.Area)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("201 - Created\n"))
}

func main(){
    http.HandleFunc("/city", postHandler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
