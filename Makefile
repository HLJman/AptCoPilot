NAME=AptCoPilot
CURRENT=${CURDIR}
DEST=/go/src/github.com/HLJman/$(NAME)
INSTANCE=13.57.189.227

default: fmt vet build

# Build
build: 
	go build -ldflags "-s" -o $(NAME)
	# polymer build --root assets

fmt: 
	go fmt $$(go list ./... | grep -v /vendor/)

vet:
	go vet $$(go list ./... | grep -v /vendor/)

run: docker_build
	docker run --rm -it -p 8000:8000 hljman/aptcopilot

docker_build:
	docker run --rm -e "CGO_ENABLED=0" -e "GOPATH=/go" -v "$(CURRENT):$(DEST)" -w "$(DEST)" golang:1.9.2 make
	docker build -t hljman/aptcopilot .

upload: docker_zip
	scp -rv -i ~/.ssh/aws aptcopilot.tgz ubuntu@$(INSTANCE):/home/ubuntu/

ssh: 
	ssh -i ~/.ssh/aws ubuntu@$(INSTANCE)

docker_zip: docker_build
	docker save hljman/aptcopilot | gzip > aptcopilot.tgz

# docker_load:
# 	gunzip -c aptcopilot.tgz | docker load

publish: docker_zip upload ssh