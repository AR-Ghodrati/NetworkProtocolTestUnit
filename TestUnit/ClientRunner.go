package TestUnit

import (
	"github.com/joho/godotenv"
	"gsm/TestUnit/Client"
)

func RunClientKCP(Count uint64) {
	_ = godotenv.Load()
	Client.RunKCP(Count)
}
func RunClientQUIC(Count uint64) {
	_ = godotenv.Load()
	Client.RunQUIC(Count)
}
