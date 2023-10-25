package collectors

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"docker-collector/pkg/config"

// 	"github.com/docker/docker/api/types"
// 	"github.com/docker/docker/client"
// )

// type DockerCli struct {
// 	cn_names []string
// 	cli      *client.Client
// }

// func (d *DockerCli) Constructor() {
// 	var err error
// 	d.cli, err = client.NewClientWithOpts(client.FromEnv)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	d.cn_names = config.LoadInitialConfig().CnNames
// }

// func (d *DockerCli) Get_stats_cn() {
// 	containers, err := d.cli.ContainerList(context.Background(), types.ContainerListOptions{})
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	for _, container := range containers {
// 		fmt.Printf("container Name: %v\n", container.Names)
// 		fmt.Printf("container ID: %v\n", container.ID)
// 	}

// }

// func TesterMK2() {
// 	cli, err := client.NewClientWithOpts(client.FromEnv)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	for _, container := range containers {
// 		fmt.Printf("container Name: %v\n", container.Names)
// 		fmt.Printf("container ID: %v\n", container.ID)
// 	}
// }
