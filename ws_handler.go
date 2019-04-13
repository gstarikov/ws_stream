package main

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net/http"
	"time"
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Printf("ws.UpgradeHTTP error -> %s\n", err) // handle error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	go func() {
		defer conn.Close()

		for {
			log.Println("lock on read...")
			msg, op, err := wsutil.ReadClientData(conn)
			log.Println("\t\t... readed")
			if err != nil {
				log.Printf("wsutil.ReadClientData error -> %s\n", err) // handle error
				w.WriteHeader(http.StatusInternalServerError)
				break
			}

			log.Printf("code[%d] msg[%s]\n", op, string(msg))

			log.Println(string(msg))
			log.Println("lock on write...")
			err = wsutil.WriteServerMessage(conn, op, msg)
			log.Println("\t\t... writed")
			if err != nil {
				log.Printf("wsutil.WriteServerMessage error -> %s\n", err) // handle error
				w.WriteHeader(http.StatusInternalServerError)
				break
			}
		}
	}()
}

func WsGenerator(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Printf("ws.UpgradeHTTP error -> %s\n", err) // handle error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	go func() {
		defer conn.Close()

		t := time.NewTicker(time.Second)
		defer t.Stop()

		for tt := range t.C {
			err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(tt.String()))
			if err != nil {
				log.Printf("wsutil.WriteServerMessage error -> %s\n", err) // handle error
				w.WriteHeader(http.StatusInternalServerError)
				break
			}
		}
	}()
}
