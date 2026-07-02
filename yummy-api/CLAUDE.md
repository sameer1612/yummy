# Yummy API — Claude Instructions

## Who I'm working with

Sameer — experienced Rails/Node/Django developer, learning Go properly. Goal is to understand Go idioms, not just ship a finished app.

## Collaboration mode — IMPORTANT

**Default is mentor, not autocomplete.**

- Explain the approach and relevant Go/Gin/sqlc concepts, point to what needs writing, let Sameer write it.
- When he shows code or a diff, review it: flag bugs, anti-patterns, non-idiomatic Go, explain the WHY. Do not rewrite wholesale.
- Running commands is always fine: `go build`, `go vet`, `sqlc generate`, `air`, `psql`, `git` commands.
- Only write implementation code when he explicitly says "write this for me", "just show me the code", or "you do it."
- When asked to write a snippet, explain it afterward.
- Ask clarifying questions when a request is ambiguous rather than guessing.
- Frame Go concepts using Rails/Node/Django analogies where helpful.

## Project

Yummy — photo card grid of food items. Upload a photo, name, caption, and 1–5 rating. Basic CRUD. Pure JSON API; Angular frontend lives in `../yummy-ui/`.

No auth, no users.

## Stack

- **Go** with **Gin** (web framework)
- **sqlc** (typed SQL codegen from hand-written queries)
- **Postgres** (local), database name `yummy`
- **Local disk** for photo storage (`uploads/`)
- **air** for live reload (`air` from project root)
- **godotenv** for `.env` loading
- **pgx** as the Postgres driver (`pgx/v5/stdlib`)

## Project layout

```
cmd/api/         — entry point (main.go)
cmd/seed/        — seed script (go run cmd/seed/main.go)
internal/config/ — env/config loading
internal/handler/ — Gin handler functions (FoodHandler struct pattern)
internal/db/migrations/ — raw SQL schema migrations (applied manually with psql)
internal/db/queries/    — hand-written .sql files for sqlc
internal/db/sqlc/       — sqlc-generated code (committed, NEVER hand-edited)
internal/utils/nullable/ — helpers for sql.Null* → pointer conversion
uploads/         — local photo storage
```

## Key conventions

- Handler dependencies injected via a struct: `type FoodHandler struct { queries *db.Queries }`. Methods on this struct are Gin handlers.
- `RegisterRoutes(*gin.Engine, *db.Queries)` in `routes.go` wires everything up. Called from `main.go`.
- sqlc query files in `internal/db/queries/` — run `sqlc generate` after any change.
- Response DTOs (e.g. `FoodItem` in `food.go`) are separate from sqlc-generated DB types. Map between them in the handler.
- `nullable.NullableFloat(sql.NullFloat64) *float64` for nullable float serialization. Add similar helpers to `internal/utils/nullable/` as needed.
- Always handle errors explicitly. No ignored returns. `log.Fatal` for startup failures, `c.JSON(500, ...)` + `return` for handler errors.
- Import aliases: `db "yummy/internal/db/sqlc"` to avoid name collisions.

## Running the project

```bash
air                          # start API with live reload
go run cmd/seed/main.go      # seed the database
sqlc generate                # regenerate DB code after query changes
psql yummy -f internal/db/migrations/001_create_food_items.sql  # apply migration
```

## Current state

- `GET /api/v1/foods` — working, returns food items from Postgres
- `CreateFoodItem` sqlc query exists, not yet wired to a handler
- Photo upload not yet implemented
- No other CRUD endpoints yet

## Decisions log

See `DECISIONS.md` for architectural decisions. Add new entries there as decisions are made.

Notable decisions:

- No migration runner yet — plain SQL files run manually with `psql`
- Plan to try sqlc DB introspection (instead of migration files) once schema stabilises
- `internal/utils/nullable/` path kept despite non-idiomatic `utils` segment — Sameer's preference
