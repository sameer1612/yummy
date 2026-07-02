# Decisions

## 003 — Photo path as plain string for now

`POST /foods` accepts `photo_path` as a plain string from the client and stores it directly. Real file upload to `uploads/` will be implemented after all CRUD endpoints are done.

## 002 — Try sqlc DB introspection later

Currently using migration files as sqlc's schema source. Plan to revisit: delete migrations and point sqlc at the live DB via introspection instead. Better fits polyglot environments where the DB schema isn't owned by a single app — same pattern as Rails' schema.rb.

## 001 — No migration runner (for now)

Skipping tools like `golang-migrate` to reduce setup friction while learning the stack. Migrations are plain SQL files run manually with `psql`. Revisit once the schema stabilizes.
