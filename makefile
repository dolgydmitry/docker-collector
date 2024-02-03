test:
	go test -v ./...
server:
	go run cmd/main/main.go

monitordeploy:
	docker-compose -f sample_monitor_deploy/docker-compose.yaml up -d

monitordelete:
	docker-compose -f sample_monitor_deploy/docker-compose.yaml down -v

