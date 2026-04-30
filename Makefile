include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d todo-postgres

env-down:
	@docker compose down todo-postgres

env-cleanup:
	@read -p "Очистить все файлы окружения valume? Опасность утери данных. [y/N]: " ans;\
	if [ "$$ans" = "y" ]; then \
		docker compose down todo-postgres port-forwarder && \
		sudo rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очитска окружения отменена";\
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm --use-aliases todo-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action=up 1"; \
		exit 1; \
	fi; \
	docker compose run --rm --use-aliases todo-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todo-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

logs-cleanup:
	@read -p "Очистить все log файлы? Опасность утери логов. [y/N]: " ans;\
	if [ "$$ans" = "y" ]; then \
		rm -rf ${PROJECT_ROOT}/logs && \
		echo "Файлы логов очищены"; \
	else \
		echo "Очитска логов отменена";\
	fi

todoapp-run: 
	@export LOGGER_FOLDER=${PROJECT_ROOT}/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go

todoapp-deploy:
	@docker compose up -d --build todoapp

todoapp-undeploy:
	@docker compose down todoapp

ps:
	@docker compose ps