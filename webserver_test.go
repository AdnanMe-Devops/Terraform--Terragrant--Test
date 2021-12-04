package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-

io/terratest/modules/http-helper"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformWebserverExample(t *testing.T) {


	// The values to pass into the Terraform CLI
	terraformOptions := terraform.WithDefaultRetryableErrors(t, 

&terraform.Options{
		
        // The path to where the example Terraform code is located
		TerraformDir: "../examples/webserver",

		// Variables to pass to the Terraform code using -

var options
		Vars: map[string]interface{}{
			"servername": "testwebserver",
		},
	})


	// Run a Terraform init and apply with the Terraform 

options
	terraform.InitAndApply(t, terraformOptions)

	// Run a Terraform Destroy at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Retrieve the public IP address using Terraform Show
	publicIp := terraform.Output(t, terraformOptions, 

"public_ip")

	// Perform an HTTP request to the Public IP to validate 

status 200 and the body contains the following string
	url := fmt.Sprintf("http://%s:8080", publicIp)
	http_helper.HttpGetWithRetry(t, url, nil, 200, "Cloud 

How are you IamGreat!", 30, 5*time.Second)
}
 
//Starting at the top package test declares that the name of the package is test. This is for describing the purpose of the package. 

Next is the import declaration. The import declaration contains libraries or packages used in the Golang code. Packages native to 

GO like fmt can be referenced as well as packages from source control repositories like GitHub. Notice some of the libraries from 

Terratest are referenced and will be used to deploy the Terraform code and perform an HTTP test to validate the web site. 
