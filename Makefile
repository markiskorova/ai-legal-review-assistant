SHELL := /bin/bash

up:
\tdocker compose up -d --build

down:
\tdocker compose down -v

logs:
\tdocker compose logs -f --tail=200

migrate:
\tdocker compose exec api ./api migrate up

seed:
\tdocker compose exec api ./api seed

db:
\tdocker compose exec postgres psql -U postgres -d legalia

fmt:
\tfind . -name '*.go' -not -path "./apps/web/*" -exec gofmt -w {} +

test:
\tgo test ./...
