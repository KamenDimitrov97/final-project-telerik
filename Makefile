GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: audit build clean delimiter fmt lint run test update help

audit: ## Audits and finds vulnerable dependencies
	go list -json -m all | nancy sleuth

build: Dockerfile ## Builds ./Dockerfile image name: project
	docker build -t project .

build-bin: Dockerfile ## Builds ./Dockerfile image name: project
	go build -o build/final-project

clean: ## Removes /bin folder
	rm -fr ./build
	rm -fr ./vendor
	
delimiter-%:
	@echo '===================${GREEN} $* ${RESET}==================='

fmt: ## Formats the code using go fmt and go vet
	go fmt ./...
	go vet ./...

lint: ## Automated checking of your source code for programmatic and stylistic errors
	golangci-lint run ./...

run: ## Run the app locally
	go run . 

run-container: build ## First builds ./Dockerfile with image name: project and then runs a container, with name: project_container, on port :28700 
	docker run -p :3333:3333 --name project_container -ti --rm project

test: ## Runs standard unit test tests
	go test -race -cover -v ./... 

update: ## Go gets all of the dependencies and downloads them
	go get .
	go mod download

help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
