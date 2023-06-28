#!/bin/bash
set -x

TEST_RESULT=0
GO_TOOLSET_IMAGE='registry.access.redhat.com/ubi9/go-toolset:1.18.9'

teardown() {

    [ "$TEARDOWN_RAN" -ne "0" ] && return

    echo "Running teardown..."

    podman rm -f "$TEST_CONTAINER_NAME"
    TEARDOWN_RAN=1
}

trap teardown EXIT ERR SIGINT SIGTERM

mkdir -p artifacts

get_N_chars_commit_hash() {

    local CHARS=${1:-7}

    git rev-parse --short="$CHARS" HEAD
}

TEST_CONTAINER_NAME="changelog-$(get_N_chars_commit_hash 7)"

podman run -d --name "$TEST_CONTAINER_NAME" \
    "$GO_TOOLSET_IMAGE" sleep infinity


podman cp -a . "$TEST_CONTAINER_NAME:/workdir"

podman exec --workdir /workdir "$TEST_CONTAINER_NAME" make install > 'artifacts/install_logs.txt'
podman exec --workdir /workdir -e PATH=/opt/app-root/src/go/bin:$PATH "$TEST_CONTAINER_NAME" make test-unit > 'artifacts/test_logs.txt'
TEST_RESULT=$?

podman cp "$TEST_CONTAINER_NAME:/workdir/cover.out" 'artifacts/cover.out'
podman cp "$TEST_CONTAINER_NAME:/workdir/junit-changelog.xml" 'artifacts/junit-changelog.xml'

if [ $TEST_RESULT -eq 0 ]; then
    echo "tests ran successfully"
else
    echo "tests failed"
    sh "exit 1"
fi