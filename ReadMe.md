## Real-Time Stock Candle Tracker
### Overview
This project is a real-time stock candle tracker that provides live updates on stock prices for multiple symbols. It consists of a Go backend server that connects to Finnhub's WebSocket API for live trade data and a mobile frontend built with Expo.
 ## Features
- Real-time stock price updates for multiple symbols (AAPL, AMZN, TSLA, GOOGL, NFLX, PYPL)

- WebSocket connection for live data streaming

- OHLC (Open, High, Low, Close) candle generation

- Historical data storage and retrieval

- Mobile app for easy access to stock information

### Backend
- Go
- Gorilla WebSocket
- GORM (with PostgreSQL driver)
- Air (for live reloading during development)
- Docker and Docker Compose

### Frontend
- React Native
- Expo
- TypeScript
- react-native-wagmi-charts (for stock charts)
- Expo Router (for navigation)

### Data Provider
- Finnhub API

## Getting Started



### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/stock-tracker.git
   cd stock-tracker
   ```

2. Install frontend dependencies:
   ```
   cd mobile
   npm install
   ```

3. Set up backend:
   ```
   cd ../backend
   go mod tidy
   ```

4. Create a `.env` file in the `backend` directory and add your Finnhub API key:
   ```
   API_KEY=your_finnhub_api_key_here
   ```

### Running the App

1. Start the backend server:
   ```
   cd backend
     make start
   ```

2. In a new terminal, start the React Native app:
   ```
   cd mobile
   npx expo start
   ```

3. Follow the Expo CLI instructions to run the app on your preferred device or emulator.