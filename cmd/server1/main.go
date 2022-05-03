package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/log"
	"os"
	"rpc-ws/server1"
)

func main() {
	port := flag.String("portNum", ":6666", "rpc server1 port number")
	flag.Parse()
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StreamHandler(os.Stderr, log.JSONFormat())))
	s := &server1.Server{}
	s.Start(*port)
}
