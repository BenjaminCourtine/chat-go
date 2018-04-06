package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
	"sync"
)

func main() {
	m := map[int]*websocket.Conn{}
	i := 0
	var mutex = &sync.Mutex{}

	http.Handle("/connws/", websocket.Handler(func(ws *websocket.Conn) {
		data := map[string]string{}
		mutex.Lock()
		m[i] = ws
		mutex.Unlock()
		i++

		for {
			err := websocket.JSON.Receive(ws, &data)
			if err != nil {
				fmt.Println(err)
				ws.Close()
				break
			}
			fmt.Println(data)
			message := fmt.Sprintf("%s à dit : %s<br/>", data["nom"], data["texte"])
			for i := range m {
				err2 := websocket.Message.Send(m[i], message)
				if err2 != nil {
					fmt.Println(m[i])
				}
			}
		}

	}))

	http.ListenAndServe(":2222", nil)

}
