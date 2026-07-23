# Project Conventions

## Stack
- Go 1.25+, Gin, GORM, PostgreSQL
- JWT auth (dgrijalva/jwt-go)
- Service layer pattern: handler → service → repository

## Code Style
- Use `log/slog` for logging (not `fmt.Println`)
- Business errors use sentinel in `services/errors.go` (`ErrInvalidInput`, `ErrInvalidStatus`, `ErrInsufficientStock`, `ErrConflict`)
- "Not found" errors use `gorm.ErrRecordNotFound` directly (not wrapped)
- Handlers use `errors.Is()` to map errors to HTTP status codes
- All async ops use `sync` or `context` — no naked goroutines

## Project Structure
- `handlers/` — HTTP request handling only
- `services/` — Business logic
- `repositories/` — GORM queries
- `models/` — Struct definitions + DTOs
- `middleware/` — JWT, logger
- `db/` — Database init
- `routes/` — Route definitions

## Response
- Success: `SuccessResponse(SuccessParams{C, Data, Message})`
- Error: `ErrorResponse(ErrorParams{C, Status, Message})`
