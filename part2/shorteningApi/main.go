package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func toBase62(url []byte) string {
	var i big.Int
	i.SetBytes(url[:])
	return i.Text(62)
}

func CreateShortUrlHandler(w http.ResponseWriter, r *http.Request){
    dbUltimate := make(map[string]string)
    requestBody := r.Body
    defer requestBody.Close()

    requestBytes, errReadAll := io.ReadAll(requestBody)
    if errReadAll != nil{
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(errReadAll.Error()))
        w.Write([]byte("\n"))
        return
    }

    shortUrl := toBase62(requestBytes)
    dbUltimate[shortUrl] = string(requestBytes)
    dbUltimateJson, errJson := json.Marshal(dbUltimate)
    if errJson != nil{
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(errJson.Error()))
        w.Write([]byte("\n"))
        return
    }

    f, errOpenFile := os.OpenFile("ultimateDB.json", os.O_CREATE | os.O_RDWR, 0755)
    if errOpenFile != nil{
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(errOpenFile.Error()))
        w.Write([]byte("\n"))
        return
    }
    defer f.Close()

    jsonFromDB, errJsonFromDb := io.ReadAll(f)
    if errJsonFromDb != nil{
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(errJsonFromDb.Error()))
        w.Write([]byte("\n"))
        return
    }

    bw := bufio.NewWriterSize(f, 16)
    defer bw.Flush()

    if len(jsonFromDB) == 0{

        _, errBufioW := bw.Write(dbUltimateJson)
        if errBufioW != nil{
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(errBufioW.Error()))
            w.Write([]byte("\n"))
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(shortUrl))
        w.Write([]byte("\n"))
        return
    }

    mapCheck := make(map[string]string)
    json.Unmarshal(jsonFromDB, &mapCheck)
    _, cond := mapCheck[string(requestBytes)]
    if  cond {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(mapCheck[string(requestBytes)]))
        w.Write([]byte("\n"))
        return
    }


    mapCheck[shortUrl] = string(requestBytes)
    fmt.Println(mapCheck)
    mapCheckJson, _ := json.Marshal(mapCheck)
    errTrunc := f.Truncate(0)
    if errTrunc != nil{
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(errTrunc.Error()))
        w.Write([]byte("\n"))
        return
    }

    _, errSeek := f.Seek(0, os.SEEK_SET)
    if errSeek != nil{
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(errSeek.Error()))
        w.Write([]byte("\n"))
        return
    }

    _, errBufioW := bw.Write(mapCheckJson)
    if errBufioW != nil{
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(errBufioW.Error()))
        w.Write([]byte("\n"))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(shortUrl))
    w.Write([]byte("\n"))
}

func RedirectOriginalHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    _, urlCond := vars["url"]
    if len(vars) > 0 && urlCond{
        f, errOpenFile := os.OpenFile("ultimateDB.json", os.O_RDONLY, 0755)
        if errOpenFile != nil{
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(errOpenFile.Error()))
            w.Write([]byte("\n"))
            return
        }
        defer f.Close()

        jsonFromDB, errJsonFromDb := io.ReadAll(f)
        if errJsonFromDb != nil{
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(errJsonFromDb.Error()))
            w.Write([]byte("\n"))
            return
        }
        checkMap := make(map[string]string)
        json.Unmarshal(jsonFromDB, &checkMap)
        if _, check := checkMap[vars["url"]]; check {
            http.Redirect(w, r, checkMap[vars["url"]], 301)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("Original Url not found"))
    w.Write([]byte("\n"))
}

func main(){
    router := mux.NewRouter()
    router.HandleFunc("/api/v1/new", CreateShortUrlHandler).Methods("POST")
    router.HandleFunc("/api/v1/{url}", RedirectOriginalHandler).Methods("GET")

    server := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    log.Fatal(server.ListenAndServe())
}

