package Client

import (
	"../models"
	"../utils"
	"crypto/sha1"
	"github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
	_ "io"
	"log"
	"os"
	"time"
)

const Msg = "There used to a be a rule that before anyone was permitted to write any code that uses TCP, they were required to repeat the following sentence from memory and explain what it means: TCP is not a message protocol, it is a reliable byte-stream protocol that does not preserve application message boundaries"

func Run(count int) {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)

	// wait for server to become ready
	time.Sleep(time.Second)
	log.Println("Connecting....")
	defer log.Println("Send", count, "Packet DONE!!")

	// dial to the echo server
	if sess, err := kcp.DialWithOptions(os.Getenv("ENDPOINT"), block, 10, 3); err == nil {
		for i := 0; i < count; i++ {
			data := time.Now().Unix()
			random, _ := utils.GenerateRandomString(200)
			_, _ = sess.Write(utils.Serialize(models.Message{SequenceNumber: uint64(i), Milis: data, Msg: random}))
			time.Sleep(time.Millisecond)
		}
	} else {
		log.Fatal(err)
	}
}
