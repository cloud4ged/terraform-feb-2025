package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"math/rand/v2"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/image"
)

func listDockerImages() {
	fmt.Println("")
	fmt.Println("List of Images in Local docker registry")
	fmt.Println("")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	images, err := cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		if len(image.RepoTags) > 0 {
			tokens := strings.Split(image.RepoTags[0],":")
			fmt.Println(tokens[0]+ ":" + tokens[1])
			//fmt.Println(image.ID)
		}
	}
	fmt.Println("")
}

func pullDockerImage( imageName string ) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
}
/*
func removeDockerImage( imageName string ) {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    image, _, err := cli.ImageInspectWithRaw(ctx, imageName)
    if err != nil {
        panic(err)
    }

    _, err = cli.ImageRemove(context.Background(), image.ID, types.ImageRemoveOptions{})
    if err != nil {
        panic(err)
    }

    fmt.Printf("Image %s removed\n", image.ID)
}
*/

func listRunningDockerContainers() {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%s %s (status: %s)\n", ctr.ID, ctr.Image, ctr.Status)
	}
}

func createNewDockerContainer( imageName string, containerName string ) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	pullDockerImage ( imageName )

	containerName = fmt.Sprintf("%v-%d",containerName, rand.IntN(99999) )

	resp, err := cli.ContainerCreate(ctx, &container.Config{Image: imageName,}, nil, nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

func main() {
	listDockerImages()
	fmt.Println()
	//listRunningDockerContainers()
	//createNewDockerContainer("bitnami/nginx:latest", "nginx")
	//createNewDockerContainer("nginx:latest", "nginx")
}
