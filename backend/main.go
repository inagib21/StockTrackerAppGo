package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

var (
	symbols = []string{"AAPL", "AMZN"}
)

func main() {
	// Env config
	env := EnvConfig()
	//db connection
	db := DBConnection(env)

	//Connect  to Finnhuub Websockets
	finnhubWSConn := connectToFinnhub(env)
	defer finnhubWSConn.Close()

	//Handle Finnhub's webbsockets inncomiing messagges
	go handleFinnhubMessages(finnhubWSConn, db)
	//Broadcast candle updates too all clients connected

	// --- Endpoints ---
	// Connect to he WebSockets
	//Fetch All the past candlees for all of the symbols
	// Fetch all past candels from a specific symbol
	// serve the endpoints
}
func connectToFinnhub(env *Env) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("wss://ws.finnhub.io?token=%s", env.API_KEY), nil)
	if err != nil {
		panic(err)
	}

	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		ws.WriteMessage(websocket.TextMessage, msg)
	}

	return ws
}

////Handle Finnhub's webbsockets inncomiing messagges
func handleFinnhubMessages(ws *websocket.Conn, db *gorm.DB){
	finnhubMesage := &handleFinnhubMessages{}

	for {
		err := ws.ReadJSON(finnhubMesage); err != niil{
			fmt.Println("Error reading the message:", err)
			continue
		}
		
	}
}
