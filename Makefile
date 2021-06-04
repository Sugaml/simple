REPOSITORY    := gcr.io/zerone-devops-labs
NAME          := 01cloud-payment
BRANCH        := $(shell git rev-parse --abbrev-ref HEAD)
HASH          := $(shell git rev-parse --short HEAD)
SUFFIX        ?= -$(subst /,-,$(BRANCH))
VERSION       := 0.1
DIRNAME       := $(shell basename $(CURDIR))
PARENTDIRNAME := $(shell basename $(shell dirname $(CURDIR)))

.PHONY: help all test format fmtcheck vet lint     qa deps clean nuke build

# Display general help about this command
help:
	@echo ""
	@echo "The following commands are available:"
	@echo ""
	@echo "    make qa          : Run all the tests"
	@echo "    make unit-test        : Run the unit tests"
	@echo ""
	@echo "    make format      : Format the source code"
	@echo "    make fmtcheck    : Check if the source code has been formatted"
	@echo "    make vet         : Check for suspicious constructs"
	@echo "    make lint        : Check for style errors"
	@echo ""
	@echo "    make deps        : Get the dependencies"
	@echo "    make clean       : Remove any build artifact"
	@echo "    make nuke        : Deletes any intermediate file"
	@echo ""

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@docker build . --target lint

.PHONY: test-dev
test-dev:
	go test ./... -cover -v

.PHONY: run-dev
run-dev:
	go run cmd/server/main.go
	
# Alias for help target
all: help

# Format the source code
format:
	@find ./ -type f -name "*.go" -exec gofmt -w {} \;

# Check if the source code has been formatted
fmtcheck:
	@mkdir -p target
	@find ./ -type f -name "*.go" -exec gofmt -d {} \; | tee target/format.diff
	@test ! -s target/format.diff || { echo "ERROR: the source code has not been formatted - please use 'make format' or 'gofmt'"; exit 1; }

# Check for syntax errors
vet:
	GOPATH=$(GOPATH) go vet ./...



# Alias to run all quality-assurance checks
qa: fmtcheck test vet lint    

# --- INSTALL ---

# Get the dependencies
deps:

# Remove any build artifact
clean:
	GOPATH=$(GOPATH) go clean ./...

# Deletes any intermediate file
nuke:
	rm -rf ./target
	GOPATH=$(GOPATH) go clean -i ./...

build:
	@echo ">> building $(REPOSITORY)/$(NAME):$(HASH)"
	@docker build -t "$(REPOSITORY)/$(NAME):$(HASH)" .
	#@docker tag "$(REPOSITORY)/$(NAME):$(VERSION)-$(DIRNAME)$(SUFFIX)" "$(REPOSITORY)/$(NAME):$(PARENTDIRNAME)-$(DIRNAME)$(SUFFIX)"
