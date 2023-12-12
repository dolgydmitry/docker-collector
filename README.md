# docker-collector
## !!!! app use DOCKER_API_VERSION=1.41 ##
## Tested only for Unix ##

### Get start 

1. build image use Dockerfile
2. specify config.yml in folder ./sample_monitor_deploy
containers:
    shall write conatner name
See sample in this directory
3. Run docker compose use make
make monitordeploy
4. Check Grafana UI on localhost:3000
admin
secret

### Desc
1. App expose container perfomance use promehteus.
2. Currently it provide two metrics:
    - cpu (percent usage)
    - memory (byte allocated)
3. It provide standart prometheus ebdpoint: /metrics


