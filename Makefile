# Copyright © 2017 Aqua Security Software Ltd. <info@aquasec.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

NAME=bucky
ORG=qualimente
PACKAGE_NAME=github.com/$(ORG)/$(NAME)
TAG=$$(git describe --abbrev=0 --tags)

LDFLAGS += -X "$(PACKAGE_NAME)/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "$(PACKAGE_NAME)/version.BuildVersion=$(shell git describe --abbrev=0 --tags)"
LDFLAGS += -X "$(PACKAGE_NAME)/version.BuildSHA=$(shell git rev-parse HEAD)"
# Strip debug information
LDFLAGS += -s

ifeq ($(OS),Windows_NT)
	suffix := .exe
endif

all: fmt vet build test

$(GOPATH)/bin/glide$(suffix):
	go get github.com/Masterminds/glide

$(GOPATH)/bin/$(NAME)$(suffix):
	go get github.com/$(ORG)/$(NAME)

glide.lock: glide.yaml $(GOPATH)/bin/glide$(suffix)
	glide update
	@touch $@

vendor: glide.lock
	glide install
	@touch $@

releases:
	mkdir -p releases

bin/linux/amd64:
	mkdir -p bin/linux/amd64

bin/darwin/amd64:
	mkdir -p bin/darwin/amd64

build: darwin linux

fmt:
	go fmt

vet:
	go vet

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

darwin: vendor releases bin/darwin/amd64
	env GOOS=darwin GOAARCH=amd64 go build -ldflags '$(LDFLAGS)' -v -o $(CURDIR)/bin/darwin/amd64/$(NAME)
	tar -cvzf releases/$(NAME)-darwin-amd64.tar.gz bin/darwin/amd64/$(NAME)

linux: vendor releases bin/linux/amd64
	env GOOS=linux GOAARCH=amd64 go build -ldflags '$(LDFLAGS)' -v -o $(CURDIR)/bin/linux/amd64/$(NAME)
	tar -cvzf releases/$(NAME)-linux-amd64.tar.gz bin/linux/amd64/$(NAME)

clean:
	rm -rf releases bin

