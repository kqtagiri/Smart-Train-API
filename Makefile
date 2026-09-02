include .env
export

run_code: 
	@go run cmd/app/main.go

run_docker:
	@docker compose up -d

run_docker_debug:
	@docker compose up

build_docker:
	@docker compose up --build

end_docker:
	@docker compose down