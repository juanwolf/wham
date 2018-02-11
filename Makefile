GO=go
PKG=./cmd

PLATFORMS=linux darwin windows
ARCH=amd64

INSTALL_GO_DEP=go get -u github.com/golang/dep/cmd/dep && dep ensure

TEST_CMD=$(GO) test -v -coverprofile=coverage.txt -covermode=atomic -race $(PKG)
BUILD_CMD=go get -u github.com/mitchellh/gox && gox -os="$(PLATFORMS)" -arch="$(ARCH)" -output="{{.Dir}}.{{.OS}}.{{.Arch}}" -ldflags "-X main.Rev=`git rev-parse --short HEAD`" -verbose ./...

DOCKER=true
DOCKER_VOLUME_PATH=/go/src/wham
DOCKER_IMAGE=wham:test

define docker_call
docker run -v $(PWD):$(DOCKER_VOLUME_PATH) $(DOCKER_IMAGE) $1;
endef

define docker_handler
	if [ "$(DOCKER)" == "true" ]; \
	then \
		docker run -v $(PWD):$(DOCKER_VOLUME_PATH) $(DOCKER_IMAGE) $1; \
	else \
		$1; \
	fi
endef


all: test build

pre:
ifeq ($(DOCKER),true)
	docker image inspect $(DOCKER_IMAGE) > /dev/null 2>&1 || 	docker build -t $(DOCKER_IMAGE) .
	$(call docker_call, $(INSTALL_GO_DEP))
else
	$(INSTALL_GO_DEP)
endif

test: pre
	$(call docker_handler, $(TEST_CMD))

build: pre
	$(call docker_handler, $(BUILD_CMD))

clean:
	rm *amd64*
