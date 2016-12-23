package lib

import (
  "fmt"
  "bufio"
  "log"
  "net"
  "os"
)

const port = "8080"

// RunHost takes an ip as argument
// and listens for connection
func RunHost(ip string) {
  ipAndPort := ip + ":" + port

  listener, listenError := net.Listen("tcp", ipAndPort)
  if listenError != nil {
    log.Fatal("Error", listenError)
  }
  fmt.Println("Listening on", ipAndPort)

  connection, acceptError := listener.Accept()
  if acceptError != nil {
    log.Fatal("Error", acceptError)
  }
  fmt.Println("New connection accepted")

  for {
    handleHost(connection)
  }
}

func handleHost(connection net.Conn) {
  reader := bufio.NewReader(connection)
  message, readError := reader.ReadString('\n')
  if readError != nil {
    log.Fatal("Error", readError)
  }
  fmt.Println("Message received: ", message)

  fmt.Print("Send message: ")
  replyReader := bufio.NewReader(os.Stdin)
  replyMessage, replyError := replyReader.ReadString('\n')
  if replyError != nil {
    log.Fatal("Error: ", replyError)
  }
  fmt.Fprint(connection, replyMessage)
}

// RunGuest takes destination ip as argument
// and connects to that ip
func RunGuest(ip string)() {
  ipAndPort := ip + ":" + port

  connection, dialError := net.Dial("tcp", ipAndPort)
  if dialError != nil {
    log.Fatal("Error: ", dialError)
  }

  for {
    handleGuest(connection)
  }
}

func handleGuest(connection net.Conn) {
  fmt.Print("Send message: ")
  reader := bufio.NewReader(os.Stdin)
  message, readError := reader.ReadString('\n')
  if readError != nil {
    log.Fatal("Error: ", readError)
  }
  fmt.Fprint(connection, message)

  replyReader := bufio.NewReader(connection)
  replyMessage, replyError := replyReader.ReadString('\n')
  if replyError != nil {
    log.Fatal("Error: ", replyError)
  }
  fmt.Println("Message received: ", replyMessage)
}
