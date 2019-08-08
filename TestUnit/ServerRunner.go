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
