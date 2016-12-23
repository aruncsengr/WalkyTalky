package main

import (
  "flag"
  "os"
  "walkytalky/lib"
)

func main() {
  var isHost bool
  flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
  flag.Parse()

  if isHost {
    // go run main.go -listen <ip>
    connectionIPconnectionIP := os.Args[2]
    lib.RunHost(connectionIPconnectionIP)
  } else {
    // go run main.go <ip>
    connectionIP := os.Args[1]
    lib.RunGuest(connectionIP)
  }

}
