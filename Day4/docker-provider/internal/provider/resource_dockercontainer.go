package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerContainer() *schema.Resource {
	return &schema.Resource{
		Description: "Manages the lifecycle of a Docker container.",

		CreateContext: resourceDockerContainerCreate,
		ReadContext:   resourceDockerContainerRead,
		UpdateContext: resourceDockerContainerUpdate,
		DeleteContext: resourceDockerContainerDelete,
		MigrateState:  resourceDockerContainerMigrateState,
		SchemaVersion: 2,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the container.",
				Required:    true,
				ForceNew:    true,
			},

			"image": {
				Type:        schema.TypeString,
				Description: "The ID of the image to back this container. The easiest way to get this value is to use the `docker_image` resource as is shown in the example.",
				Required:    true,
				ForceNew:    true,
			},

			"hostname": {
				Type:        schema.TypeString,
				Description: "Hostname of the container.",
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
			},

		},
	}
}
