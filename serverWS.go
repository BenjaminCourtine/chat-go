package main

import (
  "golang.org/x/net/websocket"
    "net/http"
    "fmt"
)

func main() {
 hub := []*websocket.Conn{}

 http.Handle("/connws/", websocket.Handler(func(ws *websocket.Conn) {
     data := map[string]string{}
     hub = append(hub, ws)

     for {
         err := websocket.JSON.Receive(ws, &data)
         if err != nil {
             fmt.Println(err)
             ws.Close()
             break
     }
     fmt.Println(data)
     message := fmt.Sprintf("%s à dit : %s<br/>", data["nom"], data["texte"])
     for i := range hub {
         websocket.Message.Send(hub[i], message)
         }
     }
 }))

 http.ListenAndServe(":2222", nil)

}





