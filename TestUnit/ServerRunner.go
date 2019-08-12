package TestUnit

import (
	"gsm/TestUnit/Server"
)

func RunServerKCP(port string) {
	Server.StartKCP(port)
}
func RunServerQUIC(port string) {
	Server.StartQUIC(port)
}

func RunServerPureTCP(port string) {
	Server.StartPureTCP(port)
}

func RunServerPureUDP(port string) {
	Server.StartPureUDP(port)
}
