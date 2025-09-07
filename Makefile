# Makefile

# Path to docker-compose dev file
DEV_COMPOSE=docker/docker-compose.dev.yml

# Start dev environment
dev-up:
	docker-compose -f $(DEV_COMPOSE) up -d --build

# Stop dev environment
dev-down:
	docker-compose -f $(DEV_COMPOSE) down

# Display dev compose logs
dev-logs:
	docker-compose -f $(DEV_COMPOSE) logs -f
