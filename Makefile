APP_VERSION := $$(shell git describe --tags --always --dirty)

build:
	docker compose --env-file .env up -d --build --build-arg APP_VERSION=$${APP_VERSION}
docker-lint:
	docker run --rm -i -v ./hadolint.yaml:/.config/hadolint.yaml hadolint/hadolint < .docker/go/Dockerfile
lint:
	docker run -t --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v2.1.6 golangci-lint run