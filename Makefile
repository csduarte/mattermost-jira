.PHONY: run build build-linux build-osx build-windows prebuild

# Golang Flags
GOPATH ?= $(GOPATH:):./vendor
GOFLAGS ?= $(GOFLAGS:)
GO=go

build: build-windows build-osx build-linux
	
build-linux: .prebuild
	env GOOS=linux GOARCH=amd64 $(GO) build -i -o $(GOPATH)/bin/linux_amd64/mattermost-jira .
	@echo Built Linux amd64 at $(GOPATH)/bin/linux_amd64/mattermost-jira

build-osx: .prebuild
	env GOOS=darwin GOARCH=amd64 $(GO) build -i -o $(GOPATH)/bin/linux_amd64/mattermost-jira .
	@echo Build OSX amd64 at $(GOPATH)/bin/mattermost-jira

build-windows: .prebuild
	env GOOS=windows GOARCH=amd64 $(GO) build -i -o $(GOPATH)/bin/linux_amd64/mattermost-jira .
	@echo Build Windows amd64 at $(GOPATH)/bin/windows_amd64/mattermost-jira

run: .prebuild
	@echo Building and Running
	$(GO) build -i -o mattermost-jira .
	./mattermost-jira

.prebuild:
	@echo Prebuild current unused 

	touch $@