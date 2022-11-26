package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode())

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func WsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Client connected to websocket endpoint")
	var resp WsJsonResponse
	resp.Message = `<em><small>Connected to server</small></em>`
	err = ws.WriteJSON(resp)
	if err != nil {
		log.Println(err)
		return
	}
}

func RenderPage(w http.ResponseWriter, templateName string, data jet.VarMap) error {
	view, err := views.GetTemplate(templateName)
	if err != nil {
		return err
	}

	return view.Execute(w, data, nil)
}
