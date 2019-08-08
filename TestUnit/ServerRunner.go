package TestUnit

import (
	"github.com/joho/godotenv"
	"gsm/TestUnit/Server"
)

func RunServerKCP() {
	_ = godotenv.Load()
	Server.StartKCP()
}
func RunServerQUIC() {
	_ = godotenv.Load()
	Server.StartQUIC()
}
