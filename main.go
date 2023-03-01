package main

import "time"

func main() {
	go startServer()
	go startClient()
	time.Sleep(time.Second * 5)
}
