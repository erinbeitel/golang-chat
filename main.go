package main

import (
  "flag"
  "fmt"
  "github.com/erinbeitel/golang-chat/Godeps/_workspace/src/github.com/gorilla/websocket"
  "log"
  "net/http"
  "errors"
  "testing"

)

var connections map[*websocket.Conn]bool

func sendAll(msg []byte) {
  for conn := range connections {
    if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
      delete(connections, conn)
      return
    }
  }
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
  //from gorilla
  conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
  if _, ok := err.(websocket.HandshakeError); ok {
    http.Error(w, "Not a websocket handshake", 400)
    return
  } else if err != nil {
    log.Println(err)
    return
  }
  defer conn.Close()
  connections[conn] = true
  for {
    _, msg, err := conn.ReadMessage()
    if err != nil {
      delete(connections, conn)
      return
    }
    log.Println(string(msg))
    sendAll(msg)
  }
}


func checkPort (port int) error {
  if port > 65535 {
    return errors.New("Port can not be greater than 65535")
  } else if port < 0 {
    return errors.New("Port can not be negative.")
  } else if port < 1024 {
    return errors.New("Ports must be between 1024 and 65535")
  } else {
    return nil
  }
}


func TestCheckPort (t *testing.T) {
	failMessage := "Failed testing of checkPort() function."

  if checkPort(-1) == nil {
    t.Errorf(failMessage)
  } else if checkPort(00000000) == nil {
    t.Errorf(failMessage)
  } else if checkPort(99999999) == nil {
    t.Errorf(failMessage)
  }
}

func checkDir (dir string) error {
  if (len(dir) == 0) {
    return errors.New("The length of the directory string was zero.")
  } else if dir == "/" {
    return errors.New("Can not run the webserver from the filesystem root. This was probably an accident.")
  } else {
    return nil
  }
}

func TestCheckDir (t *testing.T) {
  failMessage := "Failed testing of checkDir() function."

  if checkDir("/") == nil {
    t.Errorf(failMessage)
  } else if checkDir("usr/bin/test") == nil {
    t.Errorf(failMessage)
  }
}

func main() {

  port := flag.Int("port", 8000, "port to server on")
  dir := flag.String("directory", "web/", "directory of web files")
  flag.Parse()
   
  if err := checkPort(*port); err != nil {
    fmt.Println("The specified port is invalid")
    panic(err)
  }
  
  if err := checkDir(*dir); err != nil {
    fmt.Println("The specified directory is invalid.")
    panic(err)
  }
  
  connections = make(map[*websocket.Conn]bool)

  fs := http.Dir(*dir)
  fileHandler := http.FileServer(fs)
  http.Handle("/", fileHandler)
  http.HandleFunc("/ws", wsHandler)

  log.Printf("Running on port %d\n", *port)

  addr := fmt.Sprintf("127.0.0.1:%d", *port)

  err := http.ListenAndServe(addr, nil)
  fmt.Println(err.Error())
  
}
