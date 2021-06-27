APP_VERSION=v0.0.1
TAG_REVISION=$(shell git rev-parse --short HEAD)
BUILD_TAG=${APP_VERSION}-${TAG_REVISION}
DOCKER_REGISTRY=0.0.0.0:5000
ARCH?=GOARCH=amd64
PLATFORM?=linux/amd64

.PHONY: build
build:
	GOOS=linux $(ARCH) go build -o main-svc -v cmd/main.go

.PHONY: build-img
build-img: build
	docker buildx build --platform $(PLATFORM) -t main-svc:${BUILD_TAG} .

.PHONY: tag
tag: build-img
	docker tag main-svc:${BUILD_TAG} main-svc:latest

.PHONY: run
run: tag
	docker run -p 8000:80 -t main-svc:latest

.PHONY: test
test:
	go test -race ./internal/concretes/...

.PHONY: integration-test
integration-test:
	GOOS=linux $(ARCH) go build -o integration-test -v test/integration/main.go
	docker buildx build --platform $(PLATFORM) \
		-t integration-test:latest \
		-f test/integration/Dockerfile .
	cd test/integration \
		&& docker-compose up -d bureaucrat \
		&& docker-compose up tester \
		&& docker-compose down

.PHONY: publish
publish: tag
	docker tag main-svc:${BUILD_TAG} ${DOCKER_REGISTRY}/main-svc:${BUILD_TAG}
	docker tag main-svc:latest ${DOCKER_REGISTRY}/main-svc:latest
	docker push ${DOCKER_REGISTRY}/main-svc:${BUILD_TAG}
	docker push ${DOCKER_REGISTRY}/main-svc:latest

.PHONY: deploy
deploy: publish
	kubectl apply -f deploy/main-svc.yaml