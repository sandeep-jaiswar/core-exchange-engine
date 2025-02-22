package engine

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/sandeep-jaiswar/core-exchange-engine/internal/db"
	"github.com/sandeep-jaiswar/core-exchange-engine/internal/models"
)

type OrderBook struct {
    BuyOrders  []models.Order
    SellOrders []models.Order
    mu         sync.Mutex
}

func (ob *OrderBook) AddOrder(order models.Order) {
    ob.mu.Lock()
    defer ob.mu.Unlock()

    if order.Side == "buy" {
        ob.BuyOrders = append(ob.BuyOrders, order)
    } else if order.Side == "sell" {
        ob.SellOrders = append(ob.SellOrders, order)
    }

    // Persist order in MongoDB
    collection, err := db.GetCollection("orders")
    if err != nil {
        log.Printf("Failed to get MongoDB collection: %v", err)
        return
    }

    _, err = collection.InsertOne(context.Background(), order)
    if err != nil {
        log.Printf("Failed to persist order: %v", err)
    }
}

func (ob *OrderBook) MatchOrders() {
    ob.mu.Lock()
    defer ob.mu.Unlock()

    // Simple FIFO matching algorithm
    for len(ob.BuyOrders) > 0 && len(ob.SellOrders) > 0 {
        buyOrder := ob.BuyOrders[0]
        sellOrder := ob.SellOrders[0]

        if buyOrder.Price >= sellOrder.Price {
            // Execute trade
            trade := models.Trade{
                ID:           generateTradeID(),
                BuyOrderID:   buyOrder.ID,
                SellOrderID:  sellOrder.ID,
                Symbol:       buyOrder.Symbol,
                Price:        sellOrder.Price,
                Quantity:     min(buyOrder.Quantity, sellOrder.Quantity),
                ExecutedAt:   time.Now(),
            }

            // Persist trade in MongoDB
            collection, err := db.GetCollection("trades")
            if err != nil {
                log.Printf("Failed to get MongoDB collection: %v", err)
                return
            }

            _, err = collection.InsertOne(context.Background(), trade)
            if err != nil {
                log.Printf("Failed to persist trade: %v", err)
            }

            // Update order quantities
            buyOrder.Quantity -= trade.Quantity
            sellOrder.Quantity -= trade.Quantity

            // Remove fully filled orders
            if buyOrder.Quantity == 0 {
                ob.BuyOrders = ob.BuyOrders[1:]
            }
            if sellOrder.Quantity == 0 {
                ob.SellOrders = ob.SellOrders[1:]
            }
        } else {
            break // No more matches
        }
    }
}

func generateTradeID() string {
    return "trade_" + time.Now().Format("20060102150405")
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}