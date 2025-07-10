package handlers

import (
    "internal-transfers/database"
    "internal-transfers/models"
    "github.com/gofiber/fiber/v2"
)

// POST /transactions
func CreateTransaction(c *fiber.Ctx) error {
    tx := new(models.Transaction)
    if err := c.BodyParser(tx); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    db := database.DB

   
    transaction, err := db.Begin()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to start transaction",
        })
    }

  
    var sourceBalance float64
    err = transaction.QueryRow(
        "SELECT balance FROM accounts WHERE account_id = $1 FOR UPDATE",
        tx.SourceAccountID,
    ).Scan(&sourceBalance)
    if err != nil {
        transaction.Rollback()
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Source account not found",
        })
    }

    if sourceBalance < tx.Amount {
        transaction.Rollback()
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Insufficient funds",
        })
    }

  
    _, err = transaction.Exec(
        "UPDATE accounts SET balance = balance - $1 WHERE account_id = $2",
        tx.Amount, tx.SourceAccountID,
    )
    if err != nil {
        transaction.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update source account",
        })
    }

   
    _, err = transaction.Exec(
        "UPDATE accounts SET balance = balance + $1 WHERE account_id = $2",
        tx.Amount, tx.DestinationAccountID,
    )
    if err != nil {
        transaction.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update destination account",
        })
    }

   
    err = transaction.Commit()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to commit transaction",
        })
    }

    return c.SendStatus(fiber.StatusCreated)
}
