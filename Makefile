all: build

install:
	go install

build:
	go build ; cp terraform-provider-azuredevops ~/.terraform.d/plugins/terraform-provider-azuredevops; go run generate-schema/generate-schema.go && cd ~/.vim/plugged/vim-terraform-completion && ruby version_dissect.rb && cd -; rm -rf .terraform/; terraform init

.PHONY: install
