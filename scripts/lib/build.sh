# Bundle divided OpenAPI specification to single.
function build_openapi(){
  npx swagger-cli bundle \
    openapi/openapi.yaml \
    --outfile builds/openapi.yaml \
    --type yaml \
    --dereference
}

# Build client library for go
function build_client_library(){
  local RELEASE_DIR=$1
  local PACKAGE_NAME=$2

  oapi-codegen \
    -generate "types,client" \
    -package ${PACKAGE_NAME} \
    -o ${RELEASE_DIR}/main.go \
    builds/openapi.yaml
}

# Build server library for go
function build_server_library(){
  local RELEASE_DIR=$1
  local PACKAGE_NAME=$2

  oapi-codegen \
    -generate "types,server" \
    -package ${PACKAGE_NAME} \
    -o ${RELEASE_DIR}/main.go \
    builds/openapi.yaml
}

# Initialize Go Modules
function initialize_go_modules(){
  local RELEASE_DIR=$1
  local GOMODULE_NAME=$2

  # Return if go.mod exists
  if [ -f "$RELEASE_DIR/go.mod" ]; then
    return
  fi

  cd ${RELEASE_DIR} && go mod init ${GOMODULE_NAME}
  cd - > /dev/null
}

# Update Go Modules
function update_go_modules(){
  local RELEASE_DIR=$1

  cd ${RELEASE_DIR} && go mod tidy
  cd - > /dev/null
}
