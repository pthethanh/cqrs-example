REGISTRY_SERVER=192.168.98.100:5000
NAME=book-query
DOCKER_IMAGE=$(REGISTRY_SERVER)/$(NAME):latest

all: vet build redeploy

vet:
	go vet .

build:
	CGO_ENABLED=0 go build -o $(NAME).bin

format:
	gofmt -l -w .

first_deploy: build
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)
	kubectl create -f $(NAME).yaml

redeploy: build
	kubectl delete -f $(NAME).yaml
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)
	kubectl create -f $(NAME).yaml
