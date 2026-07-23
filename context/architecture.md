# Architecture

## Layer Structure

```
handler → service → repository → DB (GORM)
```

| Layer | Responsibility | Example |
|---|---|---|
| Handler | Parse HTTP request, validate input, send response | `handlers/product_handler.go` |
| Service | Business logic, orchestration, transaction | `services/product_service.go` |
| Repository | Database queries via GORM | `repositories/product_repository.go` |
| Model | Data structure, DB schema mapping | `models/product.go` |

## Flow

```
Client Request
  → Route (routes/routes.go)
    → Handler
      → Service
        → Repository
          → Database

Response
  ← Handler (SuccessResponse / ErrorResponse)
```

## Response Format

All endpoints return consistent JSON:

### Success (200)
```json
{
  "status": 200,
  "data": { ... },
  "message": "success"
}
```

### Error
```json
{
  "status": 400,
  "data": null,
  "message": "Invalid request payload"
}
```
