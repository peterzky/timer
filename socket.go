import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func SocketServer(port int, c chan string, q chan bool) {
	listen, err := net.Listen("tcp4", "127.0.0.1:"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn, c, q)
	}

}

func SocketClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		os.Exit(0)
	}

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	fmt.Println(string(buff[:n]))
}

func handler(conn net.Conn, s chan string, q chan bool) {
	defer conn.Close()
	w := bufio.NewWriter(conn)
	q <- true
	w.Write([]byte(<-s))
	w.Flush()

}
