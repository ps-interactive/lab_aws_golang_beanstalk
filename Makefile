GOCMD=go
GOTEST=$(GOCMD) test
GORUN=${GOCMD} run

all: test

start: 
				${GORUN} application.go
test:
				$(GOTEST) -v ./handlers/*
