/***
rpc 包建立在 gob 包之上
***/

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/zhangzt123/Golearn/goNET/RPC/RPCServer/rpcobjects"
)

func main() {
	calc := new(rpcobjects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
