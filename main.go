package main

import (
    "internal-transfers/database"
    "internal-transfers/handlers"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    database.Connect()

    app.Post("/accounts", handlers.CreateAccount)
    app.Get("/accounts/:account_id", handlers.GetAccount)
    app.Post("/transactions", handlers.CreateTransaction)

    app.Listen(":3000")
}
