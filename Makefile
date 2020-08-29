__DOCKER_COMPOSE=$(shell which docker-compose)
__DOCKER=$(shell which docker)
__API_SERVICE_NAME="golang-api"
__DOCKER_HUB_REPO_NAME="kazu12344"

lint:
	goreturns -w src cmd

docker-build: lint
	$(__DOCKER_COMPOSE) build \
	--build-arg SSH_KEY="`cat ~/.ssh/id_rsa`" \
	$(__API_SERVICE_NAME)

docker-run: docker-build
	$(__DOCKER_COMPOSE) down
	$(__DOCKER_COMPOSE) up

docker-push:
	$(__DOCKER) image tag $(__API_SERVICE_NAME):latest $(__DOCKER_HUB_REPO_NAME)/$(__API_SERVICE_NAME):latest
	$(__DOCKER) login
	$(__DOCKER) image push $(__DOCKER_HUB_REPO_NAME)/$(__API_SERVICE_NAME):latest

test-coverage:
	go test -race -coverprofile=cover.out -covermode=atomic ./...
	go tool cover -html=cover.out -o cover.html

clean:
	$(__DOCKER) system prune -f