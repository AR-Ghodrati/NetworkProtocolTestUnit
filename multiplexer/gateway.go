package multiplexer

import (
	"github.com/xtaci/kcp-go"
	"gsm/utils"
	"log"
)
const defaultBufferSize = uint(100000)

func Multiplex(conn *kcp.UDPSession) {
	buf := make([]byte, defaultBufferSize)
	length_r, err :=conn.Read(buf)
	log.Println(string(buf[:length_r]))
	msg := utils.Deserialize(buf[:length_r])

	log.Println(" said: ",msg)
	if err != nil {
		log.Fatal(err)
	}
}

