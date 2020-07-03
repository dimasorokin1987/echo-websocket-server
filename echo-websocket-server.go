package main

import (
  "fmt"
  //"io"
  "net/http"
  "log"
  //"golang.org/x/net/websocket"
)

indexHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  if origin := r.Header.Get("Origin"); origin != "" {
    w.Header().Set("Access-Control-Allow-Origin", origin)
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers",
        "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  }
  // Stop here if its Preflighted OPTIONS request
  if r.Method == "OPTIONS" {
    return
  }
  fmt.Fprintf(w, "Hello World!")
})

// Echo the data received on the WebSocket. 
// func EchoServer(ws *websocket.Conn) {
//   io.Copy(ws, ws)
// }

// This example demonstrates a trivial echo server.
func main() {
  //http.Handle("/echo", websocket.Handler(EchoServer))
  // err := http.ListenAndServe(":12345", nil)
  http.Handle("/", indexHandler)
  err := http.ListenAndServe(":80", nil)
  if err != nil {
    log.Fatalln("ListenAndServe: " + err.Error())
  }
}
