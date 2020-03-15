package main

import (
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
	}
	for {
		//bytes:=	make([]byte,1024)
		// reader := bufio.NewReader(os.Stdin)
		// bytes, err := reader.ReadBytes(1)
		// if err != nil {
		// 	fmt.Print("asd")
		// }
		conn.Write([]byte("asddasdd"))
	}

}
