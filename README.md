# Platform Changelog API

## Overview

The Platform Changelog is a system for keeping track of changes as they occur
across the platform by leveraging different types of notification events, such as
Github and Gitlab webhooks, as well as Tekton Pipeline tasks.

This API provides JSON responses to the requesting entity, mainly the [Platform
Changelog Frontend](https://www.github.com/redhatinsights/platform-changelog).

Github and Gitlab webhooks authenticated via secret token as described in the Git api.

## Architecture

Platform Changelog is a backend API that connects to a backend database for storing
supported incoming events. The current implementation supports a Postgres database
and respondes to incoming requests with JSON responses.

Also, we are integrating with App-SRE tooling to better capture information from RedHat tooling.

A frontend application has also been developed for displaying this information in
an easy to read, and searchable manner.

## REST API Endpoint

TODO: API Spec

## Adding A Service

To add a service to be supported by platform-changelog, follow these steps:

1. Add your tenant to `internal/config/tenant.yaml` if it is not included.
  ```yaml
  tenant-name:
    name: Tenant Name
```

2. Add the service to `internal/config/services.yaml`.
  
  ```yaml
  service-name:
    display_name: "Service Name"
    tenant: <tenant>
    gh_repo: <https://github.com/org/repo>
    branch: master # branch to be monitored
    namespace: <namespace of the project>
```

3. Submit an MR to this repo. It will be approved by an owner.

## Development

A Makefile has been provided for most common operations to get the app up and running.
A compose file is also available for standing up the service in podman.

Docker can be substituted for podman if needed.

### Prequisites

    podman
    podman-compose
    Golang >= 1.16

### Launching

    $> make -B build
    $> make run-db
    $> make run-migration
    $> make run-api DEBUG=1

Note: The `DEBUG` argument allows us to send webhooks without needing the secret token.

### Launching with a Mock Database

    $> make -B build
    $> make run-api-mock DEBUG=1

Note: This is useful to avoid having to run the database locally, but this will not persist data between runs.

The API should now be up and available on `localhost:8000`. You should be able to
see the API in action by visiting `http://localhost:8000/api/v1/services`.

### Testing POST Requests to the API Manually

Launch the API as instructed above, then we can send test requests to the API.

The app is designed to take in commit and deployment data through `/api/v1/github` and `/api/v1/tekton` respectively. Using webhooks is also included, but they will not be used to track our platform.

Test json is provided in the `tests` directory in this repo.

To send the requests, you can use curl the following makefile commands: 
- `make test-github`
- `make test-github-webhook`
- `make test-gitlab-webhook`
- `make test-tekton-task`.

From there, you should be able to open a browser and see the results populated at: http://localhost:8000/api/v1/commits. There will be commits matching the webhook data that was sent.

## Running Tests

TODO: Get some tests in here

# Get Help

This service is owned by the ConsoldeDot Pipeline team. If you have any questions, or
need support with this service, please contact the team on slack @crc-pipeline-team.

You can also raise an Issue in this repo for them to address.
