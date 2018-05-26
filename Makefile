all: | sanity install examples
.PHONY: all

sanity:
	@dep ensure
	@go list ./... | grep -v vendor/ | xargs go vet 
	@go list ./... | grep -v vendor/ | xargs golint
	@go list ./... | grep -v vendor/ | xargs go fmt
	@go list ./... | grep -v vendor/ | xargs go test --cover 
.PHONY: sanity

install:
	@go install -v github.com/KyleBanks/modoc
.PHONY: install

examples: | install
	@rm -rf ./examples/basic
	@modoc init --path ./examples/basic
	@modoc compile --source ./examples/basic --output ./examples/BASIC.md
	
	@modoc compile --source ./examples/readme --output ./README.md
.PHONY: examples

golden: | install
	@go test github.com/KyleBanks/modoc/pkg/document/markdown -args --update
.PHONY: golden
