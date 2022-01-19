#!/bin/bash

# Preset variables
REPO_NAME=logbook-go

# Load libraries
source ./scripts/lib/build.sh
source ./scripts/lib/validation.sh

# Validations
validate_current_dir ${REPO_NAME}
validate_npx_exists
validate_swagger_cli_exists
validate_oapi_codegen_exists

# Builds
build_openapi
build_library
