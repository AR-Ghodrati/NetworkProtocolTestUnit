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

func Run() {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)

	// wait for server to become ready
	time.Sleep(time.Second)
	log.Println("Connecting....")

	// dial to the echo server
	if sess, err := kcp.DialWithOptions(os.Getenv("ENDPOINT"), block, 10, 3); err == nil {
		for {
			data := time.Now().Unix()
			_, _ = sess.Write(utils.Serialize(models.Message{Milis: data, Msg: "dsfsfsdf"}))
			time.Sleep(time.Second)
		}
	} else {
		log.Fatal(err)
	}
}
