package multiplexer

import (
	"../utils"
	"github.com/xtaci/kcp-go"
	"log"
	"math"
	"time"
)

const defaultBufferSize = uint(16384)

func Multiplex(conn *kcp.UDPSession, count uint64) {
	buf := make([]byte, defaultBufferSize)
	//map [time Diff] = Count
	var DiffMap = make(map[float64]uint64)
	defer func() {
		for diff, Count := range DiffMap {
			log.Println("With", diff, "Millisecond -> ", "Count is:", Count)
		}
	}()

	for {
		lengthR, err := conn.Read(buf)
		msg := utils.Deserialize(buf[:lengthR])
		i := time.Now().Unix()

		DiffMap[math.Abs(float64(i-msg.Milis))] = DiffMap[math.Abs(float64(i-msg.Milis))] + 1

		if err != nil {
			log.Fatal(err)
		}
		if msg.SequenceNumber == count-1 {
			break
		}
	}
}
