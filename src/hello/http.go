package main
 
import (
    "log"
    "net/http"
)
 
func main() {
    h := http.FileServer(http.Dir("/Users/eddie_lee/go/src/http/"))
    err := http.ListenAndServe(":9999", h)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
