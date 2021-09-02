# Contributing

This document contains instructions and best practices for developing against `terraform-provider-hivelocity`.

**NOTE: This project is still in Alpha, development best practices are still being refined**

## Assumptions

1. We assume you are developing on OSX.  This document will be updated to support Microsoft and Linux distros in the future.
2. We assume you are working against our local development version of the core api.  If you are not a Hivelocity employee, replace `http://localhost:5065` with `https://core.hivelocity.net`. Be aware you will be charged when using `core.hivelocity.net`. If you'd like to improve the provider without worrying about being charged, please open an issue and leave us an email address to contact you at. 
3. You are familiar with the basics of using Terraform to create custom API calls: https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers
4. You are familiar with the best practices for extending Terraform: https://www.terraform.io/docs/extend/best-practices/index.html

## ENV VARS

You may set `HIVELOCITY_API_URL` to `http://localhost:5065/api/v2` for development.
You may set `HIVELOCITY_API_KEY` if you'd like.

## Building the most recent hivelocity-client-go

If API endpoints have been added to the core, you will want to update the go client.
To rebuild the client and add it to your project repo:

1. Install https://github.com/swagger-api/swagger-codegen
2. run the following from the folder `terraform-provider-hivelocity`:

`java -jar ~/Downloads/swagger-codegen-cli-2.4.15.jar generate -i http://localhost:5065/api/v2/swagger.json?partner=1 -l go -o ./hivelocity-client-go`

## Generating the documentation for the registry

The documentation for the Hashicorp registry is composed of the README.md file, along with the files present in the examples and docs folder.

If there were changes in the provider, the documentation will have to be updated. This can be accomplished using the following:

`go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs`

## Rebuilding the project

https://www.terraform.io/upgrade-guides/0-13.html#new-filesystem-layout-for-local-copies-of-providers

Whenever you update the project run the following cmd to add the changes to your Terraform plugins:

`go build -o ~/.terraform.d/plugins/registry.terraform.io/hivelocity/hivelocity/0.1.0/darwin_amd64/terraform-provider-hivelocity`

Watch out for the naming convention (OS and distribution).

It is recommend you set the environment variable `export TF_LOG=DEBUG` so that you can see and debug API calls while developing new functionality.

## Data Sources

All data sources should be added to the `examples/data-sources/` folder for development.

Once you have updated the example, you can test your new data source from the root of the repo:

`go build -o ~/.terraform.d/plugins/registry.terraform.io/hivelocity/hivelocity/0.1.0/darwin_amd64/terraform-provider-hivelocity && cd examples/data-sources/<data-source> && terraform init -plugin-dir ~/.terraform.d/plugins/ && terraform apply --auto-approve`

## Resources

All resources should be added to the `examples/resources/` folder for development.

Once you have updated the example, you can test your new data source from the root of the repo:

To Create/Update:

`go build -o ~/.terraform.d/plugins/registry.terraform.io/hivelocity/hivelocity/0.1.0/darwin_amd64/terraform-provider-hivelocity && cd examples/resources/hivelocity_bare_metal_device && terraform init -plugin-dir ~/.terraform.d/plugins/ && terraform apply --auto-approve`

To Delete:

`go build -o ~/.terraform.d/plugins/registry.terraform.io/hivelocity/hivelocity/0.1.0/darwin_amd64/terraform-provider-hivelocity && cd examples/resources/hivelocity_bare_metal_device && terraform init -plugin-dir ~/.terraform.d/plugins/ && terraform delete --auto-approve`

Testing provider code
---------------------------

We will mostly acceptance test the provider. So you should not run them all unless you want to create tons of resources.
If you make a change, when you are satisfied with the results, run the one covering whatever you changed.

NOTE: The acceptance test run will create real resources, and potentially charge you if you are developing against the prod API `core.hivelocity.net`.
Please open an issue and share your email address if you would like to run tests without being charged.

Before running tests set `export TF_ACC=1`

Then find the relevant test function in `*_test.go` and run 

`go test -v -timeout=10m -run=YOUR_TEST_NAME`

You can debug your API requests byl setting your environment variable `TF_LOG=DEBUG`

Testing the provider with Terraform
---------------------------------------

Once you've built the plugin binary (see [Developing the provider](#developing-the-provider) above), it can be incorporated within your Terraform environment using the `-plugin-dir` option. Subsequent runs of Terraform will then use the plugin from your development environment.

```sh
$ terraform init -plugin-dir $GOPATH/bin
```
