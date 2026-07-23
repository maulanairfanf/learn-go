# Error Standard

## Error Types

| Error | HTTP Status | Usage |
|---|---|---|
| `gorm.ErrRecordNotFound` | 404 | Record not found (DB query returns empty) |
| `services.ErrInvalidInput` | 400 | Invalid request data (wrong password, invalid category IDs) |
| `services.ErrInvalidStatus` | 400 | Order status transition not allowed |
| `services.ErrInsufficientStock` | 400 | Not enough stock for order |
| `services.ErrConflict` | 409 | Duplicate data (username already taken) |
| Other errors | 500 | Internal server error |

## Handler Pattern

```go
order, err := orderService.Pay(id)
if err != nil {
    switch {
    case errors.Is(err, gorm.ErrRecordNotFound):
        ErrorResponse(ErrorParams{C: c, Status: 404, Message: "Order not found"})
    case errors.Is(err, services.ErrInvalidStatus):
        ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Order cannot be paid"})
    default:
        ErrorResponse(ErrorParams{C: c, Status: 500, Message: "Internal server error"})
    }
    return
}
```

## Rules

1. "Not found" errors use `gorm.ErrRecordNotFound` directly (not wrapped)
2. Business logic errors use sentinel errors from `services/errors.go`
3. Repository methods return GORM errors directly
4. Handlers use `errors.Is()` to differentiate 404 / 400 / 409 / 500
