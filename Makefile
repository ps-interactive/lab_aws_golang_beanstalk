GOCMD=go
GOTEST=$(GOCMD) test
GORUN=${GOCMD} run

all: test

start: 
				${GORUN} main.go
test:
				$(GOTEST) -v ./handlers/*