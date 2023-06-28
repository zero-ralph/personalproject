GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run
GOTIDY=$(GOMOD) tidy
APP=../main.go
BUILD_OUTPUT=../../packages/build/dist/recipeapp

all: clean test build

tidy:
	$(GOTIDY)
	
clean:
	$(GOCLEAN)

dep:
	$(GOMOD) tidy
	$(GOMOD) vendor

build:
	$(GOBUILD) -o $(BUILD_OUTPUT) $(APP)

run:
	$(GORUN) $(APP)
