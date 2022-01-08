# Kantin Project

Connecting seller and buyer. Well, sort of

## NOTES

Mostly still in development. *main branch will be subject of force-push*

### Phase 1

- [ ] Refactor storage
- [x] Use database pool
- [ ] Middleware

### Phase 2

- [ ] Split handler
- [ ] Initial work on Firebase
- [ ] Refactor connection

### Phase 3

- [ ] Use Redis (JWT, probably notification ?)
- [ ] Write test

## Run

### Docker

```bash
docker-compose up -d
```

### Manual Run

Need to run Postgres instance separately on port 5432

```bash
go run cmd/server.go
```
