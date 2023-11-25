# Use for lite start
PROJECT = go-exp

go-run:
	docker exec -it $(PROJECT)_go-exp-service_1 bash -c "go run ./cmd/app/main.go"

go-build:
	docker exec -it $(PROJECT)_go-exp-service_1 bash -c "go build -o /go/bin/app ./cmd/app/main.go"; docker-compose -f docker-compose.yml -f docker-compose-ep.yml -p $(PROJECT) up
	
d-up:
	docker-compose -p $(PROJECT) up 

d-stop:	
	docker-compose -p $(PROJECT) stop

d-build:
	docker-compose -p $(PROJECT) up --build

d-log-up:
	@docker-compose -f docker-compose-graylog.yml -p $(PROJECT) up

d-metrics-up:
	@docker-compose -f docker-compose-grafana.yml -p $(PROJECT) up

kill-SIGINT:
	@docker-compose kill -s SIGINT go-exp-service
	
