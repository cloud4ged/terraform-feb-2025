package provider

import (
	"context"
	"log"
	"io"
	"os"
	"fmt"
	"math/rand/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/image"
)

func resourceDockerContainer() *schema.Resource {
	return &schema.Resource{
		Description: "Manages the lifecycle of a Docker container.",

		CreateContext: resourceDockerContainerCreate,
		ReadContext:   resourceDockerContainerRead,
		UpdateContext: resourceDockerContainerUpdate,
		DeleteContext: resourceDockerContainerDelete,

		Schema: map[string]*schema.Schema{
			"container_name": {
				Type:        schema.TypeString,
				Description: "The name of the container.",
				Required:    true,
				ForceNew:    true,
			},

			"image_name": {
				Type:        schema.TypeString,
				Description: "Name of the docker image",
				Required:    true,
				ForceNew:    true,
			},

			"host_name": {
				Type:        schema.TypeString,
				Description: "Hostname of the container.",
				Optional:    true,
			},
		},
	}
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

func resourceDockerContainerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { 
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	imageName 	  := d.Get("image_name").(string)
	containerName     := d.Get("container_name").(string)
	//containerHostname := d.Get("host_name").(string)

	pullDockerImage ( imageName )

	containerName = fmt.Sprintf("%v-%d",containerName, rand.IntN(100) )

	resp, err := cli.ContainerCreate(ctx, &container.Config{Image: imageName}, nil, nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	d.Set("container_name", containerName)
	d.SetId(resp.ID)

	return resourceDockerContainerRead(ctx, d, meta) 
}


func resourceDockerContainerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { 
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containerName     := d.Get("container_name").(string)

	removeOptions := container.RemoveOptions{RemoveVolumes: true, Force: true}

	if err := cli.ContainerRemove(ctx, containerName, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return nil 
	}

	return nil
}

func resourceDockerContainerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { 
	//Here, we need to list the running containers, find the container matching the container name that we created
	//Extract the ID, name, image and hostname properties and assign it correctly
	//In the interest of time, I have hard coded these values
//	d.SetId("123")
//	d.Set("name", "nginx") 
//	d.Set("image", "bitnami/nginx:latest")
//	d.Set("hostname", "nginx")

	return nil
}

func resourceDockerContainerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { 
	return nil
}
