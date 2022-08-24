package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const addr = "0.0.0.0:12345"
const proto = "tcp4"

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	//количество подключившихся клиентов
	counter := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		counter += 1
		fmt.Println("count", counter)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		proverbs := []string{
			"Don't communicate by sharing memory, share memory by communicating.",
			"Concurrency is not parallelism.",
			"Channels orchestrate; mutexes serialize.",
			"The bigger the interface, the weaker the abstraction.",
			"Make the zero value useful.",
			"interface{} says nothing.",
			"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
			"A little copying is better than a little dependency.",
			"Syscall must always be guarded with build tags.",
			"Cgo must always be guarded with build tags.",
			"Cgo is not Go.",
			"With the unsafe package there are no guarantees.",
			"Clear is better than clever.",
			"Reflection is never clear.",
			"Errors are values.",
			"Don't just check errors, handle them gracefully.",
			"Design the architecture, name the components, document the details.",
			"Documentation is for users.",
			"Don't panic.",
		}

		conn.Write([]byte(proverbs[rand.Intn(len(proverbs))] + "\n"))
		time.Sleep(3 * time.Second)
	}
}
