package main

import (
	"bufio"
	"fmt"
	"gsm/TestUnit"
	"log"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select Server Type : ")
	fmt.Println("1-QUIC")
	fmt.Println("2-KCP")

	for {
		char, _, _ := reader.ReadRune()
		switch char {
		case '1':
			TestUnit.RunServerQUIC()
			break
		case '2':
			TestUnit.RunServerKCP()
			break
		default:
			log.Println("Select Valid Option")
		}
	}

}
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
