package TestUnit

import (
	"github.com/joho/godotenv"
	"gsm/TestUnit/Client"
)

func RunClient() {
	_ = godotenv.Load()
	Client.Run(Count)
}
