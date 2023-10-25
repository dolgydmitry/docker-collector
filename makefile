addenv:
	export DOCKER_API_VERSION=1.41
test:
	go test ./... -v
run:
	go run cmd/main/main.go
showpid:
	lsof -t -i :8091

promdeploy:
	docker run \
	-p 9090:9090  \
	-d \
	--name prom_test \
	--add-host=host.docker.internal:host-gateway \
	-v /Users/A118582519/go/src/golang_learn/aiflow_exporter/prom_config/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus 

promkill:
	docker stop prom_test
	docker rm prom_test


monitordeploy:
	docker-compose -f sample_monitor_deploy/docker-compose.yaml up -d

monitordelete:
	docker-compose -f sample_monitor_deploy/docker-compose.yaml down -v