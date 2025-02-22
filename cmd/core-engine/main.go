package main

import (
	"time"

	"github.com/sandeep-jaiswar/core-exchange-engine/config"
	"github.com/sandeep-jaiswar/core-exchange-engine/internal/db"
	"github.com/sandeep-jaiswar/core-exchange-engine/internal/engine"
	"github.com/sandeep-jaiswar/core-exchange-engine/internal/models"
)

func main() {
    config.InitConfig()

    db.InitMongoDB()

    orderBook := engine.OrderBook{}

    orderBook.AddOrder(models.Order{
        ID:        "1",
        UserID:    "user1",
        Symbol:    "AAPL",
        Side:      "buy",
        Price:     150.0,
        Quantity:  100,
        Status:    "open",
        CreatedAt: time.Now(),
    })

    orderBook.AddOrder(models.Order{
        ID:        "2",
        UserID:    "user2",
        Symbol:    "AAPL",
        Side:      "sell",
        Price:     149.0,
        Quantity:  50,
        Status:    "open",
        CreatedAt: time.Now(),
    })

    orderBook.MatchOrders()
}