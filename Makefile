DEBUG=0

POSTGRES_PASSWORD=crc
POSTGRES_USER=crc
POSTGRES_DB=gumbaroo
POSTGRES_PORT=5432:5432
POSTGRES_HOST=localhost

GITHUB_WEBHOOK_KEY = "1234567890"
GITHUB_WEBHOOK_SIGNATURE := $(shell cat tests/github_webhook.json | openssl dgst -sha256 -hmac "$(GITHUB_WEBHOOK_KEY)" | sed -e 's/.*= //')

GITLAB_WEBHOOK_KEY = "0987654321"

.PHONY: build

build: platform-changelog-api platform-changelog-migration

platform-changelog-api:
	go build -o $@ cmd/api/main.go

platform-changelog-migration:
	go build -o $@ internal/migration/main.go

lint:

	gofmt -l .
	gofmt -s -w .

test:

	go test -p 1 -v ./...

run-migration: platform-changelog-migration

	./platform-changelog-migration

run-api: platform-changelog-api

	GITHUB_WEBHOOK_SECRET_TOKEN=$(GITHUB_WEBHOOK_KEY) GITLAB_WEBHOOK_SECRET_TOKEN=$(GITLAB_WEBHOOK_KEY) ./platform-changelog-api

run-api-mock: platform-changelog-api

	GITHUB_WEBHOOK_SECRET_TOKEN=$(GITHUB_WEBHOOK_KEY) GITLAB_WEBHOOK_SECRET_TOKEN=$(GITLAB_WEBHOOK_KEY) DB_IMPL=mock ./platform-changelog-api

run-db:

	podman run --rm -it -p ${POSTGRES_PORT} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_DB=${POSTGRES_DB} --name postgres postgres:12.4

check-db:

	psql -h ${POSTGRES_HOST} --user ${POSTGRES_USER} --db ${POSTGRES_DB}

test-github-webhook:

	curl -X POST -H "X-Hub-Signature-256: sha256=$(GITHUB_WEBHOOK_SIGNATURE)" -H "X-Github-Event: push"   -H "Content-Type: application/json" --data-binary "@tests/github_webhook.json" http://localhost:8000/api/platform-changelog/v1/github-webhook

test-gitlab-webhook:

	curl -X POST -H "X-Gitlab-Token: $(GITLAB_WEBHOOK_KEY)" -H "X-Gitlab-Event: Push Hook" -H "Content-Type: application/json" --data "@tests/gitlab_webhook.json" http://localhost:8000/api/platform-changelog/v1/gitlab-webhook

test-tekton-task:

	curl -X POST http://localhost:8000/api/platform-changelog/v1/tekton --data "@tests/tekton/valid.json" -H "Content-Type: application/json"

compose:

	podman-compose -f development/compose.yml up

compose-quiet:

	podman-compose -f development/compose.yml up -d

compose-down:

	podman-compose -f development/compose.yml down

clean:
	go clean
	rm -f platform-changelog-api platform-changelog-migration
