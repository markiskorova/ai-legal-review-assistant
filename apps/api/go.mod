module github.com/markiskorova/ai-legal-review-assistant/apps/api

go 1.24.3

require (
	github.com/go-chi/chi/v5 v5.0.11
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/jackc/pgx/v5 v5.5.5
	github.com/joho/godotenv v1.5.1
	github.com/markiskorova/ai-legal-review-assistant/pkg v0.0.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/markiskorova/ai-legal-review-assistant/pkg => ../../pkg
