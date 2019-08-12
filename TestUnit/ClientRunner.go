package TestUnit

import (
	"gsm/TestUnit/Client"
)

func RunClientKCP(address string, Count uint64, packetStringSize int) {
	Client.RunKCP(address, Count, packetStringSize)
}
func RunClientQUIC(address string, Count uint64, packetStringSize int) {
	Client.RunQUIC(address, Count, packetStringSize)
}

func RunClientTCP(address string, Count uint64, packetStringSize int) {
	Client.RunPureTCP(address, Count, packetStringSize)
}

func RunClientUDP(address string, Count uint64, packetStringSize int) {
	Client.RunPureUDP(address, Count, packetStringSize)
}
