package main

const MAKEFILE = \
`
WINDOWS = GOOS=windows GOARCH=amd64 CGO_ENABLED=0
LINUX = GOARCH=amd64 CGO_ENABLED=0

SRCDIR = src/{{.Repo}}
OUTPUT = -o /go/Release/{{.Name}}
OPTS = -a -v

all: get windows linux

get: ;\
go get -d -v {{.Repo}}

windows: ;\
cd $(SRCDIR); \
$(WINDOWS) go build $(OUTPUT).exe $(OPTS) ./...

linux: ;\
go clean; \
pwd; \
cd $(SRCDIR); \
$(LINUX) go build $(OUTPUT) $(OPTS) ./...
`
