#!/bin/bash

# Preset variables
REPO_NAME=logbook-openapi-src
GIT_LATEST_TAG=$(git describe --tags --abbrev=0)

# Load libraries
source ./scripts/lib/build.sh
source ./scripts/lib/release.sh
source ./scripts/lib/validation.sh

# Validations
validate_current_dir ${REPO_NAME}
validate_npx_exists
validate_swagger_cli_exists
validate_oapi_codegen_exists
validate_goreleaser_exists

# Builds
build_openapi

# Build client
GO_CLIENT_RELEASE_DIR=releases/client
release_library ${GO_CLIENT_RELEASE_DIR} ${GIT_LATEST_TAG}

# Release server
GO_SERVER_RELEASE_DIR=releases/server
release_library ${GO_SERVER_RELEASE_DIR} ${GIT_LATEST_TAG}
