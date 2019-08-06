package main

import (
	"../gsm/Client"
	"../gsm/protocol"
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

const Count = 100

func main() {
	_ = godotenv.Load()

	go Client.Run(Count)
	//go Client.Run(Count)
	//go Client.Run(Count)

	protocol.StartKCP(Count)

}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}
