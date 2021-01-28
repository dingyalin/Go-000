package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Message ...
type Message struct {
	data  []byte
	write *bufio.Writer
	mux   sync.Mutex
}

var wg sync.WaitGroup
var signChan = make(chan os.Signal, 1)
var connChan = make(chan net.Conn)
var messageChan = make(chan *Message, 10)

func listenSign() {
	signal.Notify(signChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}

func accept(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error: " + err.Error())
		} else {
			connChan <- conn
		}
	}
}

func handleConn() {
	for {
		select {
		case conn := <-connChan:
			remoteAddr := conn.RemoteAddr()
			log.Println("Client " + remoteAddr.String() + " connected")
			func(conn net.Conn) {
				defer conn.Close()
				// 读写缓冲区
				rd := bufio.NewReader(conn)
				wr := bufio.NewWriter(conn)
				mux := sync.Mutex{}
				for {
					line, _, err := rd.ReadLine()
					if err != nil {
						if err.Error() == "EOF" {
							log.Printf("%s close connection\n", remoteAddr)
							return
						}

						log.Printf("%s read error: %v\n", remoteAddr, err)
						return
					}
					messageChan <- &Message{data: line, write: wr, mux: mux}

				}

			}(conn)
		}
	}
}

func handMessage() {
	for {
		select {
		case message := <-messageChan:
			data := message.data
			wg.Add(1)
			message.mux.Lock()

			// 业务逻辑
			time.Sleep(100 * time.Millisecond)

			message.write.WriteString("hello ")
			message.write.Write(data)
			message.write.Flush() // 一次性syscall

			message.mux.Unlock()
			wg.Done()
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	log.Println("Listening ...")

	go listenSign()
	go accept(listener)
	go handleConn()
	go handMessage()

	select {
	case <-signChan:
		wg.Wait()
		log.Fatalf("exit")
	}

}
