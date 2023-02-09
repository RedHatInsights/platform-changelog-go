DEBUG=0

POSTGRES_PASSWORD=crc
POSTGRES_USER=crc
POSTGRES_DB=gumbaroo
POSTGRES_PORT=5432:5432
POSTGRES_HOST=localhost

build:

	go build -o platform-changelog-api cmd/api/main.go

lint:

	gofmt -l .
	gofmt -s -w .

test:

	ginkgo -r -v

run-migration:

	go build -o platform-changelog-migration internal/migration/main.go
	./platform-changelog-migration

run-api:

	DEBUG=${DEBUG} ./platform-changelog-api

run-api-mock:
	DEBUG=${DEBUG} DB_IMPL=mock ./platform-changelog-api

run-db:

	podman run --rm -it -p ${POSTGRES_PORT} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_DB=${POSTGRES_DB} --name postgres postgres:12.4

check-db:

	psql -h ${POSTGRES_HOST} --user ${POSTGRES_USER} --db ${POSTGRES_DB}

test-github-webhook:

	curl -X POST -H "X-Github-Event: push" -H "Content-Type: application/json" --data "@tests/github_webhook.json" http://localhost:8000/api/platform-changelog/v1/github-webhook

test-gitlab-webhook:

	curl -X POST -H "X-Gitlab-Event: Push Hook" -H "Content-Type: application/json" --data "@tests/gitlab_webhook.json" http://localhost:8000/api/platform-changelog/v1/gitlab-webhook

test-tekton-task:

	curl -X POST http://localhost:8000/api/platform-changelog/v1/tekton --data "@tests/tekton/valid.json" -H "Content-Type: application/json"

compose:

	podman-compose -f development/compose.yml up

compose-quiet:

	podman-compose -f development/compose.yml up -d

compose-down:

	podman-compose -f development/compose.yml down
