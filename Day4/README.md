# Day 4

## Lab - Developing a custom Terraform docker provider (Incomplete exercise - do not try this )
Let's install the Terraform Plugin Framework
```
mkdir terraform-provider-docker
cd terraform-provider-docker
go mod init github.com/tektutor/terraform-provider-docker
go get github.com/hashicorp/terraform-plugin-framework
go get github.com/docker/docker/client
```

Add the below code into the provider.go file
<pre>
package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// DockerProvider implements the Terraform provider
type DockerProvider struct{}

func (p *DockerProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "docker"
}

func (p *DockerProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *DockerProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDockerContainerResource,
	}
}

func NewDockerProvider() provider.Provider {
	return &DockerProvider{}
}
</pre>

Add the below code in resource_docker_container.go
<pre>
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/attribute"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DockerContainerResource represents a Terraform resource for managing Docker containers
type DockerContainerResource struct{}

func NewDockerContainerResource() resource.Resource {
	return &DockerContainerResource{}
}

func (r *DockerContainerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "tektutor_docker_container"
}

func (r *DockerContainerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": attribute.StringAttribute{
				Required: true,
			},
			"image": attribute.StringAttribute{
				Required: true,
			},
			"id": attribute.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DockerContainerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data struct {
		Name  types.String `tfsdk:"name"`
		Image types.String `tfsdk:"image"`
		ID    types.String `tfsdk:"id"`
	}

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		resp.Diagnostics.AddError("Docker Client Error", err.Error())
		return
	}

	log.Printf("Creating Docker container: %s with image %s", data.Name.ValueString(), data.Image.ValueString())

	containerConfig := &container.Config{
		Image: data.Image.ValueString(),
	}

	containerBody, err := cli.ContainerCreate(ctx, containerConfig, nil, nil, nil, data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Container Creation Failed", err.Error())
		return
	}

	if err := cli.ContainerStart(ctx, containerBody.ID, types.ContainerStartOptions{}); err != nil {
		resp.Diagnostics.AddError("Failed to Start Container", err.Error())
		return
	}

	data.ID = types.StringValue(containerBody.ID)

	resp.State.Set(ctx, &data)
}

func (r *DockerContainerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data struct {
		ID types.String `tfsdk:"id"`
	}

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		resp.Diagnostics.AddError("Docker Client Error", err.Error())
		return
	}

	_, err = cli.ContainerInspect(ctx, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Container Not Found", err.Error())
		resp.State.RemoveResource(ctx)
		return
	}
}

func (r *DockerContainerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {}

func (r *DockerContainerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data struct {
		ID types.String `tfsdk:"id"`
	}

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		resp.Diagnostics.AddError("Docker Client Error", err.Error())
		return
	}

	log.Printf("Removing Docker container: %s", data.ID.ValueString())

	err = cli.ContainerRemove(ctx, data.ID.ValueString(), types.ContainerRemoveOptions{Force: true})
	if err != nil {
		resp.Diagnostics.AddError("Container Deletion Failed", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}	
</pre>

Add the below code in main.go
<pre>
package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	providerserver.Serve(context.Background(), NewDockerProvider, providerserver.ServeOpts{
		Address: "registry.terraform.io/tektutor/docker",
	})
}


</pre>

Build and Test
```
go build -o terraform-provider-docker
```

Move the binary to Terraform plugin directory
```
mkdir -p ~/.terraform.d/plugins/docker/docker/1.0.0/linux_amd64
mv terraform-provider-docker ~/.terraform.d/plugins/tektutor/docker/1.0.0/linux_amd64/
```

Create a main.tf with below content
<pre>
terraform {
  required_providers {
    example = {
      source  = "tektutor/docker"
      version = "1.0.0"
    }
  }
}

provider "tektutor-docker" {}

resource "tektutor_docker_container" "my_container" {
  name  = "my-container"
  image = "nginx"
}
</pre>

Execute Terraform automation scripts
```
terraform init
terraform apply
docker ps
```
