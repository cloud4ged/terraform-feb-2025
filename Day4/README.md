# Day 4

## Lab - Developing a custom Terraform docker provider (Incomplete exercise - do not try this )
Let's install the Terraform Plugin Framework
```
go mod init github.com/tektutor/terraform-provider-docker
go get github.com/hashicorp/terraform-plugin-framework
```

Create the terraform provider directory structure
```
mkdir terraform-provider-docker
cd terraform-provider-docker
touch main.go provider.go resource_tektutor_docker.go
```

Add the below code into the provider.go file
<pre>
package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Provider implementation
type DockerProvider struct{}

func (p *DockerProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tektutor-docker"
}

func (p *DockerProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *DockerProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDockerResource,
	}
}

func NewDockerProvider() provider.Provider {
	return &DockerProvider{}
}  
</pre>

Add the below code in resource_tektutor_docker.go
<pre>
package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/attribute"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Define the resource
type DockerResource struct{}

func NewDockerResource() resource.Resource {
	return &ExampleResource{}
}

func (r *DockerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "tektutor_dockercontainer_resource"
}

func (r *DockerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": attribute.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *DockerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data struct {
		Name types.String `tfsdk:"name"`
	}

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.Set(ctx, &data)
}

func (r *DockerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {}

func (r *DockerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {}

func (r *DockerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {}  
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

resource "tektutor_dockercontainer_resource" "nginx_container" {
  imageName = "bitnami/nginx:latest"
  containerName = "nginx1"
}
</pre>

Execute Terraform automation scripts
```
terraform init
terraform apply
```
