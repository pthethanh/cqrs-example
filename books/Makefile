REGISTRY_SERVER=192.168.98.100:5000
NAME=books
DOCKER_IMAGE=$(REGISTRY_SERVER)/$(NAME):latest

all: vet format build first_deploy

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

re_deploy: build
	kubectl delete deployment,service $(NAME)
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)
	kubectl create -f $(NAME).yaml
