package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Handler struct {
	Service Service
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		func() {
			time.Sleep(time.Second * 3)
			memInfo, err := h.Service.MemoryUtilization()
			if err != nil {
				log.Println(err)
			}
			h.SendMessage(conn, 1, memInfo)

			time.Sleep(time.Second * 3)
			cpuInfo, err := h.Service.CpuUtilization()
			if err != nil {
				log.Println(err)
			}

			h.SendMessage(conn, 1, cpuInfo)

			time.Sleep(time.Second * 3)
			diskInfo, err := h.Service.DiskUtilization()
			if err != nil {
				log.Println(err)
			}

			h.SendMessage(conn, 1, diskInfo)
		}()
	}
}

func (h *Handler) SendMessage(conn *websocket.Conn, msgType int, payload any) {
	msgData, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	if err := conn.WriteMessage(msgType, msgData); err != nil {
		log.Println(err)
		return
	}
}
