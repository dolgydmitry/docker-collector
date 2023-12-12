addenv:
	export DOCKER_API_VERSION=1.41
test:
	go test ./... -v
run:
	go run cmd/main/main.go
showpid:
	lsof -t -i :8099

monitordeploy:
	docker-compose -f sample_monitor_deploy/docker-compose.yaml up -d

monitordelete:
	docker-compose -f sample_monitor_deploy/docker-compose.yaml down -v


pprof-heap:
	go tool pprof -inuse_objects http://localhost:6060/debug/pprof/heap

pprof-alloc-inuse:
	go tool pprof  -inuse_space  http://localhost:6060/debug/pprof/allocs

pprof-alloc:
	go tool pprof http://localhost:6060/debug/pprof/allocs


pprof-make-trace:
	curl -o trace.out http://localhost:6060/debug/pprof/trace\?seconds\=180

pprof-view-trace:
	go tool trace trace.out