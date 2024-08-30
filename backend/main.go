package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	symbols = []string{"AAPL", "AMZN"}

	// Broadcast messages to all connected clients
	Broadcast = make(chan *BroadcastMessaage)

	//	Map all connected clients and symbol they're subscribed to
	clientConns = make(map[websocket.Conn]string)

	// Map of all ongoing live candles for each symbol
	tempCandles = make(map[string]*TempCandle)

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
	go broadcastUpdates()

	// --- Endpoints ---
	// Connect to he WebSockets
	http.HandleFunc("/ws",WSHandler)
	//Fetch All the past candlees for all of the symbols
	// Fetch all past candels from a specific symbol
	// serve the endpoints
}

//webb socket enpoint to connect clients to the latest updates on the symbbol theyre subscribes too
func WSHandler(w http.ResponseWriter,r *http.Request) {
upgrader := websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool{
		return true
	} ,
}
conn, err := upgrader.Upgrade(w, r, nil)
if err != nil{
	log.Println("Failed to Upgrade connection: %e", err)
}
// CLose ws connection & unregister the client whenn they disconnect
defer conn.Close()
defer func(){
	delete(clientConns, conn)
	log.Println("Client disconnected")
}()

//Register the new client to the symbol theyre subscribing to 
for {
	_, symbol, err := conn.ReadMessaage()
	clientConns(conn) = string(symbol)
	log.Println("New Client Conneccted!")

	if err != nil {
		log.Println("error reading from the client:", err)
		break
	}
}
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
	finnhubMesage := &FinnhubMessages{}

	for {
		if err := ws.ReadJSON(finnhubMesage); err != nil{
			fmt.Println("Error reading the message:", err)
			continue
		}
		// Only try to  process the message data if its a trade operation
		if finnHubMessage.Type == "trade"{
			for _, := finnfinnhubMesage.Data{
				//Process the trade data
				procecessTradeData(trade, db)

			}
		}

	}
}


//Proccess each trrade and update or creaate temporrary candles

func processTradeData(trade *TradeData, db *gorm.DB) {
	//Protect the goroutine from data rraces

	mu.Lock()
	defer mu.Unlock()

	// Extract trade data
	symbol := trade.Symbol
	price := trade.Price
	volume := float64(trade.Volume)
	timestamp := time.UnixMillii(trade.Timestamp)

	//Retrieve or create a tempCandel for the symbol
	tempCandle, exists := tempCadles[symbol]

	// if the  trmmpCandels does noot exist or hsoould be already closed
	if  !exists || timestramp.After(tempCandle.CloseTime) {
		// Finalize and sabbe the previous candle, start a new on
		if exists {
			//convert the tempCandle to a Candle
			candle := tempCandle.toCandle()

			// Save candle to the db
			if err := db.Create(candle).Error; err != nil {
				fmt.Println("Error saving the candle to the DB:", err)
			}
			//Broadcast the closed candle 
			broadccast <- &BroadcastMessaage{
				UpdateType: Closed,
				Candle: candle,
			}
		}
	
	//  Initialize a new candle

	tempCandle = &tempCandle{
		Symbol: symbol,
		OpenTime: timestamp,
		CloseTime: timestamp.Add(time.Minute),
		OpenPrice: price,
		ClosePrice: price,
		HighPrice: price,
		LowPrice:price , 
		Volume: tempCandle.Volume,

	}
}
//update current tempCandle with new trade data
tempCandle.ClosePrice = price
tempCandle.Volume += volume
if price < tempCandle.HighPrice{
	tempCandle.HighPrice = price
}
if price < tempCandle.LowPrice{
	tempCandle.LowPrice = price
}
// Store the Tempcandle for the symbol
tempCandles[symbol] = tempCandle

//Write to the broadcast channel live ongoing channel
broadcast <- &BroadcastMessaage{
	UpdateType: Live,
	Candle: tempCandle.toCandle(),
}


}

// Send candle updates to client conneccted every 1 sexond at maximum, unlesss its a closed candle
func broadcasUpdates(){
	// Set the broadcast interval to 1 second
	ticker  := time.NewTicker(1 * time.Second)
	defer ticket.Stop()

	var latestUpdate *BroadcastMessaage

	for {
		select{
			// Watch for new uupdatees from the broadcast channel
		case update := <- Broadcast:
			//if the update is a closed candle, broadcast it immediately
			if update.UpdateType == Closed {
				//broadcasst it
				broadCastToClients(latestUpdate)
			} else {
				// replace temp updates
				latestUpdate = update
			}
		case <- ticket.C:
			// Broadcast the latest update
			if latestUpdate != nil{
				//Broadcast it 
				broadCastToClients(latestUpdate)
			}
			latestUpdate = nil
		}
	}
}
//  Broadcast updates to clients
func broadCastToClient(update *BroadcasMessage){
	//Marshall the updatee struct into json
	jsonUpdate, _ := json.Marshal(update)

	//send the uupdate to all connected clients subscribes tot the symbool
	for clientConn, symbol := range clientConns {
		// if the client is subscribed to the cymmbe of the update
		if update.Candle.Symbol == symbbol {
			// Send thee update to the cliient 
			err := clientCnon.WriteMessage(websocket.TextMessage, jsonUpdate)
			if err != nil{
				log.Println("Error sending message to client:", err)
				clientConn.Close()
				delete(clientConns, clientConn)
			}
		}
	}
}