up:
	if docker network ls --format "{{.Name}}" | grep user_echo_network; then \
		echo "Network 'user_echo_network' already exists"; \
	else \
		echo "Creating network 'user_echo_network'..."; \
		docker network create user_echo_network; \
	fi

	docker compose build
	docker compose down
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down
	docker volume prune -a -f

prune:
	make stop
	docker compose down
	docker system prune -a -f
	docker volume prune -a -f

format:
	go fmt ./...

run:
	go run cmd/app/main.go
