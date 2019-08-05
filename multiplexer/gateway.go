package multiplexer

import (
	"../utils"
	"github.com/xtaci/kcp-go"
	"log"
	"math"
	"time"
)

const defaultBufferSize = uint(100000)

func Multiplex(conn *kcp.UDPSession) {
	buf := make([]byte, defaultBufferSize)
	lengthR, err := conn.Read(buf)
	//log.Println(string(buf[:lengthR]))
	msg := utils.Deserialize(buf[:lengthR])

	i := time.Now().Unix()

	log.Println("Client Time : ", msg.Milis, " , Server Time : ", i, " , Deff : ", math.Abs(float64(i-msg.Milis)))
	if err != nil {
		log.Fatal(err)
	}
}
