.PHONY: up down logs ps dbshell seed reset

up:
	docker compose -f infra/docker-compose.yml up -d --build

down:
	docker compose -f infra/docker-compose.yml down -v

logs:
	docker compose -f infra/docker-compose.yml logs -f

ps:
	docker compose -f infra/docker-compose.yml ps

dbshell:
	docker exec -it alra_db psql -U postgres -d legal_assistant

seed:
	# Optional: run a seed SQL after migrations
	docker exec -i alra_db psql -U postgres -d legal_assistant < apps/api/migrations/0002_seed_playbook.sql

reset: down up
