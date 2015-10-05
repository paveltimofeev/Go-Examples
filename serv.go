package main

import "net/http"
import "time"
import "log"

func main() {
 s := &http.Server{
    Addr:           ":8080",
    //Handler:        myHandler,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }  
  log.Fatal(s.ListenAndServe())
}
