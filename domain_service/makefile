GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run
GOTIDY=$(GOMOD) tidy
APP=./main.go
APPNAME=domainService
BUILD_OUTPUT=build/dist/$(APPNAME)

all: clean test build

tidy:
	$(GOTIDY)
	
clean:
	$(GOCLEAN)

dep:
	$(GOMOD) tidy
	$(GOMOD) vendor

service:
	$(GOBUILD) -o $(BUILD_OUTPUT) $(APP)

run:
	$(GORUN) $(APP)

