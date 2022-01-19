#!/bin/bash

# Preset variables
REPO_NAME=logbook-openapi-src

# Load libraries
source ./scripts/lib/build.sh
source ./scripts/lib/validation.sh

# Validations
validate_current_dir ${REPO_NAME}
validate_npx_exists
validate_swagger_cli_exists
validate_oapi_codegen_exists

# Build OpenAPI
build_openapi

# Build client
GO_CLIENT_PACKAGE_NAME=logbook
GO_CLIENT_MODULE_NAME=github.com/umatare5/logbook-go
GO_CLIENT_RELEASE_DIR=releases/client

validate_goreleaser_yml_exists ${GO_CLIENT_RELEASE_DIR}
build_client_library  ${GO_CLIENT_RELEASE_DIR} ${GO_CLIENT_PACKAGE_NAME}
initialize_go_modules ${GO_CLIENT_RELEASE_DIR} ${GO_CLIENT_MODULE_NAME}
update_go_modules     ${GO_CLIENT_RELEASE_DIR}

# Build server
GO_SERVER_PACKAGE_NAME=framework
GO_SERVER_MODULE_NAME=github.com/umatare5/logbook-api-framework-impl
GO_SERVER_RELEASE_DIR=releases/server

validate_goreleaser_yml_exists ${GO_SERVER_RELEASE_DIR}
build_server_library  ${GO_SERVER_RELEASE_DIR} ${GO_SERVER_PACKAGE_NAME}
initialize_go_modules ${GO_SERVER_RELEASE_DIR} ${GO_SERVER_MODULE_NAME}
update_go_modules     ${GO_SERVER_RELEASE_DIR}
