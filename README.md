# Terraform--Terragrant--Test
used Terratest to write a test in GO for validating a Terraform   module. 
will write a test for a Terraform module using Terratest in GO.
One common misconception is that because Terraform is declarative, writing tests is a waste of time. Terraform will produce an error 
message if the resource doesn't deploy, so why write a test for it? Tests should be meaningful. Don't write tests to validate a 

Terraform resource setting. This will quickly become difficult to maintain and provides very little value. Instead, write tests to 
validate conditional logic, variable outcomes, and infrastructure functionality.
This test can now be executed every time after modifying the Terraform module to confirm that it works as expected quickly. An even better approach would be to store the module code in a Git repository and add this test to a CI pipeline, so changes will automatically be validated before merging. 

Note: Terratest should be used in a separate Azure subscription from production. It's recommended to have an Azure subscription just for testing Terraform code. This way automated cleanup can be run against the account to ensure resources are not left behind.
