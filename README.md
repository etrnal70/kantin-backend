# Kantin Project

Connecting seller and buyer. Well, sort of

## NOTES

Mostly still in development

### Phase 1

- [ ] Refactor
- [x] Use database pool
- [ ] Middleware
- [ ] API Documentation
- [ ] Automate test data

### Phase 2

- [x] Split handler
- [ ] Initial work on Firebase
- [ ] Refactor connection
- [ ] Implement proper logging
- [ ] Proper documentation
- [ ] Move to DEFAULT repo

### Phase 3

- [ ] Use Redis (JWT, probably notification ?)
- [ ] Write test
- [ ] Presentation (out of project scope probably)
- [ ] CONTRIBUTING.md

## Run

### Docker

```bash
docker-compose up -d
```

### Manual Run

Need to run Postgres instance separately on `:5432`

```bash
go run cmd/server.go
```
