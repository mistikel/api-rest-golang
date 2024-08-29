package main

import (
	"mezink/server"
)

func main() {
	s := server.Init()
	s.Serve()
}
