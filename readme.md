# Internal Transfers API

## Overview
A simple Go HTTP service for internal transfers between accounts.

## Requirements
- Go
- PostgreSQL

## Setup
1. Update `.env` with your database credentials.
2. Run the SQL to create the `accounts` table.
3. Build & run:
   ```bash
   go mod tidy
   go run main.go
