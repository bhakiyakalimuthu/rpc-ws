package main

import (
	"github.com/ethereum/go-ethereum/log"
	"os"
	"rpc-ws/server"
)

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StreamHandler(os.Stderr, log.JSONFormat())))
	s := &server.Server{}
	s.Start()
}
