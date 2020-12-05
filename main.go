package main

import (
	"fmt"
	"github.com/Gumkle/consoler/consoler"
	"net"
)

type Host struct {
	address net.TCPAddr
	alias string
}

func main() {
	logger := consoler.NewLogger()
	self := setupChat()
	hosts := setupHosts(self, logger)
	go listenForHostChanges(&hosts)
	go beginChat(&hosts)
}

func setupChat() *Host {
	usernameRequest := "Username: %s"
	var username string
	fmt.Printf(usernameRequest)
	fmt.Scanf(usernameRequest, &username)
	ipAddr := net.Addr().String()
	host := new(Host)
	host.address = ipAddr
	host.alias = username
	return host
}
