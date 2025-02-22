package models

import "time"

type Order struct {
    ID        string    `bson:"_id"`
    UserID    string    `bson:"user_id"`
    Symbol    string    `bson:"symbol"`
    Side      string    `bson:"side"`
    Price     float64   `bson:"price"`
    Quantity  int       `bson:"quantity"`
    Status    string    `bson:"status"`
    CreatedAt time.Time `bson:"created_at"`
}

type Trade struct {
    ID           string    `bson:"_id"`
    BuyOrderID   string    `bson:"buy_order_id"`
    SellOrderID  string    `bson:"sell_order_id"`
    Symbol       string    `bson:"symbol"`
    Price        float64   `bson:"price"`
    Quantity     int       `bson:"quantity"`
    ExecutedAt   time.Time `bson:"executed_at"`
}

type MarketData struct {
    Symbol    string      `bson:"symbol"`
    Bids      []OrderBook `bson:"bids"`
    Asks      []OrderBook `bson:"asks"`
    LastTrade Trade       `bson:"last_trade"`
    Timestamp time.Time   `bson:"timestamp"`
}

type OrderBook struct {
    Price    float64 `bson:"price"`
    Quantity int     `bson:"quantity"`
}