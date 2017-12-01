NAME=AptCoPilot
CURRENT=${CURDIR}
DEST=/go/src/github.com/HLJman/$(NAME)

default: fmt vet build

# Build
build: 
	go build -ldflags "-s" -o $(NAME)

fmt: 
	go fmt $$(go list ./... | grep -v /vendor/)

vet:
	go vet $$(go list ./... | grep -v /vendor/)

run: build
	./$(NAME)

docker_build:
	docker run --rm -e "CGO_ENABLED=0" -e "GOPATH=/go" -v "$(CURRENT):$(DEST)" -w "$(DEST)" golang:1.9.2 make
	# docker build -t $(REGISTRY):$(BRANCH) .

upload: docker_build
	scp -r -P 2222 ./AptCoPilot ./assets/* copilotadmin@50.87.144.12:/home1/copilotadmin/public_html

