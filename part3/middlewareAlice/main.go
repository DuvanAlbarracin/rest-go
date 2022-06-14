package main

import (
    "github.com/justinas/alice"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
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

func filterContentTyep(handler http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Currently checking the content of the middleware")
        if r.Header.Get("Content-Type") != "application/json"{
            w.WriteHeader(http.StatusMethodNotAllowed)
            w.Write([]byte("415 - Method Not Allowed"))
            return
        }
        handler.ServeHTTP(w, r)
    })
}

func setServerTimeCookie(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
        http.SetCookie(w, &cookie)
        log.Println("Currently setting the server time middleware")
        handler.ServeHTTP(w, r)
    })
}

func main(){
    originalHandler := http.HandlerFunc(postHandler)
    chain := alice.New(filterContentTyep, setServerTimeCookie).Then(originalHandler)
    //http.Handle("/city", filterContentTyep(setServerTimeCookie(originalHandler)))
    http.Handle("/city", chain)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
