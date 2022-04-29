package server

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

var upgrader = websocket.Upgrader{}

type Server struct{}

func (s *Server) Start() {
	log.Info("start server......")
	http.HandleFunc("/", s.handleRPC)
	http.HandleFunc("/ws/", s.handleWS)
	if err := http.ListenAndServe(":6666", nil); err != nil {
		log.Crit("failed to listen and serve", "err", err)
	}
}
func (s *Server) handleRPC(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("failed to read all", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	req := new(JsonRpcRequest)
	if err = json.Unmarshal(bodyBytes, req); err != nil {
		log.Error("failed to unmarshal", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info("request received", "request", req)
	if req.Method == "eth_blockNumber" {
		res := JsonRpcResponse{
			Jsonrpc: "2.0",
			ID:      1,
			Result:  "0xe005f7",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	if req.Method == "eth_gasPrice" {
		res := JsonRpcResponse{
			Jsonrpc: "2.0",
			ID:      1,
			Result:  "0x10267542ab",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
}

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("failed to read all", "err", err)
		return
	}
	defer conn.Close()

	for {
		req := new(JsonRpcRequest)
		if err = conn.ReadJSON(req); err != nil {
			log.Error("failed to read ws connection", "err", err)
			return
		}
		log.Info("wss: request received", "request", req)
		if req.Method == "eth_blockNumber" {
			res := JsonRpcResponse{
				Jsonrpc: "2.0",
				ID:      1,
				Result:  "0xe005f7",
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		if req.Method == "eth_gasPrice" {
			res := JsonRpcResponse{
				Jsonrpc: "2.0",
				ID:      1,
				Result:  "0x10267542ab",
			}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
}

type JsonRpcRequest struct {
	Id      interface{}   `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Version string        `json:"jsonrpc,omitempty"`
}

type JsonRpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}
