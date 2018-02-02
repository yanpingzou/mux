package main

import "efk/src/server"

func main() {
	stopCh := make(chan struct{})

	server.Init()

	<-stopCh
}
