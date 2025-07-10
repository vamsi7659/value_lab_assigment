package handlers

import (
    "internal-transfers/database"
    "internal-transfers/models"
    "github.com/gofiber/fiber/v2"
)


func CreateAccount(c *fiber.Ctx) error {
    account := new(models.Account)

    if err := c.BodyParser(account); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    _, err := database.DB.Exec(
        "INSERT INTO accounts (account_id, balance) VALUES ($1, $2)",
        account.AccountID, account.InitialBalance,
    )
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create account",
        })
    }

    return c.SendStatus(fiber.StatusCreated)
}


func GetAccount(c *fiber.Ctx) error {
    accountID := c.Params("account_id")
    var account models.Account

    row := database.DB.QueryRow(
        "SELECT account_id, balance FROM accounts WHERE account_id = $1", accountID,
    )
    err := row.Scan(&account.AccountID, &account.Balance)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Account not found",
        })
    }

    return c.JSON(account)
}
