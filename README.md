## docker-collector
 !!!! app use DOCKER_API_VERSION=1.41 ##
 Tested only for Unix ##

### Get start 
###-------------------------------------------------------------------------
1. clone repo
2. change folder permision for grafana chmod 777 -R sample_monitor_deploy/grafana/
3. specify in file config.yml in folder ./sample_monitor_deploy/collector_cfg
   
    containers:
        shall write conatner name
    See sample in this directory
   
4. Run docker compose use make:
    make monitordeploy
   
6. Check Grafana UI on localhost:3000
admin
secret

Sample dashboard located in folder "Docker collector"

### Description
###-------------------------------------------------------------------------
1. App expose containers perfomance use promehteus.
2. User can specify which containers shall be observe. For this purpose use file config.yml:

for more example, see get start topic, point 2

4. Currently it provide two metrics:
    - cpu (percent usage)
    - memory (byte allocated)
5. App expose calculated values, ready to use by human (cpu usage in percetn and memory utilization in byte).
6. It provide standart prometheus endpoint: /metrics


