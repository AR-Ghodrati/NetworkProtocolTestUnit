package TestUnit

import (
	"../TestUnit/Server"
)

const Count = 100

func RunServer() {
	Server.StartKCP(Count)
}
