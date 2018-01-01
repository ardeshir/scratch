package main

// go build -ldflags "-X main.version=0.0.1"

import (
   v "github.com/ardeshir/version"
)

var version string

func main() {

  v.V(version)
   
}
