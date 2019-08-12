package Server

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/lucas-clemente/quic-go"
	"github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
	"gsm/Multiplexer"
	"log"
	"math/big"
	"net"
	"os"
)

func StartPureTCP(port string) {

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	}
	log.Println("Pure TCO Server:\tSUCCESS On " + listener.Addr().String())

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("ERROR", err)
			continue
		}
		log.Println("Accept New Client With IP:", conn.LocalAddr().String())
		go Multiplexer.MultiplexPureTCP(conn)
	}

}

func StartPureUDP(port string) {

	listener, err := net.Listen("udp", "localhost:"+port)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	}
	log.Println("Pure UDP Server:\tSUCCESS On " + listener.Addr().String())

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("ERROR", err)
			continue
		}
		log.Println("Accept New Client With IP:", conn.LocalAddr().String())
		go Multiplexer.MultiplexPureUDP(conn)
	}

}

func StartKCP(port string) {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	if listener, err := kcp.ListenWithOptions("localhost:"+port, block, 10, 3); err == nil {
		log.Println("KCP Server:\tSUCCESS On " + listener.Addr().String())

		defer listener.Close()
		for {
			s, err := listener.AcceptKCP()
			log.Println("Accept New Client With IP:", s.LocalAddr().String())

			if err != nil {
				log.Fatal(err)
			}
			go Multiplexer.MultiplexKCP(s)
		}
	} else {
		log.Fatal(err)
	}
}

func StartQUIC(port string) {
	listener, err := quic.ListenAddr("localhost:"+port, generateTLSConfig(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QUIC Server:\tSUCCESS On " + listener.Addr().String())
	defer listener.Close()
	for {
		sess, err := listener.Accept(context.Background())
		log.Println("Accept New Client With IP:", sess.LocalAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		go Multiplexer.MultiplexQUIC(sess)
	}
}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
