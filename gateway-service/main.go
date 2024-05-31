package main

import "gateway/server"

func main() {
	httpServer := server.InitHttpServer()
	httpServer.Start()
}
