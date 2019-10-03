all: build

install:
	go install

build:
	go build ; cp terraform-provider-azuredevops ~/.terraform.d/plugins/terraform-provider-azuredevops

.PHONY: install
