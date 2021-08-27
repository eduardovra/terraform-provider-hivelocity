ifeq ($(HIVELOCITY_API_URL),)
	HIVELOCITY_API_URL:=http://localhost:5065/api/v2
	#HIVELOCITY_API_URL:=https://core.hivelocity.net
endif

ifeq ($(GOPATH),)
	GOPATH:=$(shell go env GOPATH)
endif

ifeq ($(GOOS),)
	GOOS:=$(shell go env GOHOSTOS)
endif

ifeq ($(GOARCH),)
	GOARCH:=$(shell go env GOHOSTARCH)
endif

OSARCH:=$(GOOS)_$(GOARCH)
BUILDPATH:=~/.terraform.d/plugins/registry.terraform.io/hivelocity/hivelocity/0.1.0/$(OSARCH)
SWAGGER_CODEGEN:=https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.15/swagger-codegen-cli-2.4.15.jar

default: build

install: build
	go install

build: client
	go build -o $(BUILDPATH)/terraform-provider-hivelocity

swagger-codegen-cli.jar:
	curl -o swagger-codegen-cli.jar $(SWAGGER_CODEGEN)

client: swagger-codegen-cli.jar
	rm -rf hivelocity-client-go
	java -jar swagger-codegen-cli.jar generate -i swagger.json -l go -o ./hivelocity-client-go

docs: build
	rm -rf docs
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

swagger:
	curl -o swagger.json $(HIVELOCITY_API_URL)/swagger.json?partner=1
