package main

import (
    "fmt"
    "io"
    "log"
    "net/http"

    "golang.org/x/net/websocket"
)

const listenAddr = "localhost:4000"

func main() {
    http.HandleFunc("/", rootHandler)
    http.Handle("/socket", websocket.Handler(socketHandler))
    err := http.ListenAndServe(listenAddr, nil)
    if err != nil {
        log.Fatal(err)
    }
}
