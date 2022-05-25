package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func execCommand(command string, arguments ...string) string{
    cmd, err := exec.Command(command, arguments...).Output()
    if err != nil{
        panic(err)
    }
    return string(cmd)
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params){
    //whereGo := execCommand("/usr/bin/which", "go")
    //response := execCommand(string(whereGo), "version")
    response := execCommand("/usr/local/go/bin/go", "version")
    fmt.Fprintf(w, "%s", response)
    return
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params){
    response := execCommand("/bin/cat", params.ByName("name"))
    fmt.Fprintf(w, "%s", response)
    return
}

func main(){
    router := httprouter.New()
    router.GET("/api/v1/go-version", goVersion)
    router.GET("/api/v1/show-file/:name", getFileContent)
    log.Fatal(http.ListenAndServe(":8000", router))
}
