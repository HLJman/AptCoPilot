NAME=AptCoPilot
CURRENT=${CURDIR}
DEST=/go/src/github.com/HLJman/$(NAME)
INSTANCE=13.57.189.227
REPO=joshadambell43/aptcopilot
VERSION=0.0.1

default: fmt vet build

# Build
build:
	go build -ldflags "-s" -o $(NAME)

fmt: 
	go fmt $$(go list ./... | grep -v /vendor/)

vet:
	go vet $$(go list ./... | grep -v /vendor/)

run: docker_build
	docker run --rm -it \
	-p 8000:8000 \
	$(REPO):$(VERSION)
	# -e DB_USERNAME=jadmin \
	# -e DB_PASSWORD=/etc/db_password \
	# -e DB_SERVER=database:8889 \
	# -e DB_NAME=copilot \

ssh: 
	ssh -i ~/.ssh/aws ubuntu@$(INSTANCE)

assets_build:
	cd assets; polymer build

docker_build:
	docker run --rm -e "CGO_ENABLED=0" -e "GOPATH=/go" -v "$(CURRENT):$(DEST)" -w "$(DEST)" golang:1.9.2 make
	docker build -t $(REPO):$(VERSION) .

docker_push: assets_build docker_build
	docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_PASSWORD}
	docker push $(REPO):$(VERSION)

publish: docker_push
	ssh -i ~/.ssh/aws ubuntu@$(INSTANCE) "export DOCKER_HUB_REPO=$(REPO):$(VERSION) && export DOCKER_HUB_USERNAME=${DOCKER_HUB_USERNAME} && export DOCKER_HUB_PASSWORD=${DOCKER_HUB_PASSWORD} && ./run"