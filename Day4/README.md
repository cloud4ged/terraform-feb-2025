# Day 4

## Lab - Develop a custom Terraform provider that manages local file

We need to follow a specific folder structure as recommended by the latest Terraform Provider Plugin developemnt framework
```
cd ~
mkdir -p terraform-provider-file
cd terraform-provider-file
touch go.mod
touch go.sum
touch main.go
mkdir -p internal/provider
touch internal/provider/provider.go
touch internal/provider/resource_localfile.go
tree
```

Modifiy the go.mod with the below content
<pre>
module github.com/tektutor/terraform-provider-file

go 1.22

require github.com/hashicorp/terraform-plugin-sdk/v2 v2.25.0

require (
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320 // indirect
	github.com/hashicorp/go-hclog v1.4.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.8 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.16.1 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.14.3 // indirect
	github.com/hashicorp/terraform-plugin-log v0.8.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.1.0 // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.1 // indirect
	github.com/zclconf/go-cty v1.12.1 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200711021454-869866162049 // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)	
</pre>

Modify the main.go with the below content
<pre>
package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/tektutor/terraform-provider-file/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: provider.New(version)}

	if debugMode {
		debugOpts := &plugin.ServeOpts{ProviderFunc: provider.New(version), ProviderAddr: "registry.terraform.io/tektutor/file", Debug: true}
		plugin.Serve(debugOpts)
		return
	}

	plugin.Serve(opts)
}	
</pre>

Modify internal/provider/provider.go with below content
<pre>
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	fmt.Println("New func invoked ...")
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				//"localfile": dataSourceLocalFile(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"localfile": resourceLocalFile(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type FileConfig struct {
	name string
	content string
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
		fileConfig := FileConfig {}
		fileConfig.name = "./test.txt"
		fileConfig.content = "This is a test file created by terraform"

		return &fileConfig, nil
	}
}	
</pre>

Modify internal/resource_localfile.go with below content
<pre>
package provider

import (
        "fmt"
        "os"
	"io/ioutil"
	"log"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLocalFile() *schema.Resource {
	return &schema.Resource{
		Description: "This resource will support create, reading, updating and delete a file resource via Terraform.",

		CreateContext: resourceCreateFile,
		ReadContext:   resourceReadFile,
		UpdateContext: resourceUpdateFile,
		DeleteContext: resourceDeleteFile,

		Schema: map[string]*schema.Schema{
			"file_name": {
				Description: "Name of the file.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_content": {
				Description: "Content that must be stored/retrieved to/from the file.",
				Type: 	     schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceCreateFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//client := meta.(*FileConfig)
	fmt.Println("Inside resourceCreateFile")

	//the file_name and file_content values provided in the main.tf is retrieved here
	fileName := d.Get("file_name").(string)
	content  := d.Get("file_content").(string)

	myfile, err := os.Create( fileName )

	if err != nil {
	   panic(err)
	}

	bytesWritten, err := myfile.WriteString( content + "\n")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote %d bytes into the file\n", bytesWritten)
	myfile.Sync()

	return nil 
}

func resourceReadFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
//	client := meta.(*FileConfig)

	fileName := d.Get("file_name").(string)

	content, err := ioutil.ReadFile(fileName)
	fmt.Println(content)
        if err != nil {
		log.Fatal(err)
	}
	//d.Set("content") = content 

	return nil 
}

func resourceUpdateFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil 
}

func resourceDeleteFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil 
}	
</pre>

