//go:generate npx cdktf-cli get

package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cardil/cdk-sample-curl/generated/anschoewe/curl/datacurl"
	curlprovider "github.com/cardil/cdk-sample-curl/generated/anschoewe/curl/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here
	curlprovider.NewCurlProvider(stack, jsii.String("curl"),
		&curlprovider.CurlProviderConfig{})

	latestCurl := datacurl.NewDataCurl(stack, jsii.String("kn-event-artifacts"),
		&datacurl.DataCurlConfig{
			HttpMethod: jsii.String("GET"),
			Uri:        jsii.String("https://api.github.com/repos/knative-sandbox/kn-plugin-event/releases/latest"),
		})

	latest := cdktf.NewTerraformLocal(stack, jsii.String("latest"),
		cdktf.Fn_Jsondecode(latestCurl.Response()),
	)

	m := latest.AsAnyMap()
	latestTag := cdktf.Fn_Lookup(m, jsii.String("tag_name"), jsii.String("unknown"))

	cdktf.NewTerraformOutput(stack, jsii.String("latestTag"), &cdktf.TerraformOutputConfig{
		Value: latestTag,
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdk-embedding")

	app.Synth()
}
