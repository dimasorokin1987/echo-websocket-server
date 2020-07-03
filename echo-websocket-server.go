package main

import (
 "fmt"
 "io"
 "net/http"
 "log"
 "golang.org/x/net/websocket"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w, "Hello World!")
}

// Echo the data received on the WebSocket. 
func EchoServer(ws *websocket.Conn) {
 io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func main() {
 http.Handle("/echo", websocket.Handler(EchoServer))
 // err := http.ListenAndServe(":12345", nil)
 http.HandleFunc("/",indexHandler)
 http.ListenAndServe(":80", nil)
 if err != nil {
  log.Fatalln("ListenAndServe: " + err.Error())
 }
}
