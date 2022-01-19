# Check current directory
function validate_current_dir(){
  local EXPECT_NAME=$1
  local CURRENT_DIR=$(basename "$PWD")
  if [ "${CURRENT_DIR}" != "${EXPECT_NAME}" ]; then
    echo -e "\nInvalid working directory. Please move to '${EXPECT_NAME}'.\n"
    exit 1
  fi
}

# Check npx exists
function validate_npx_exists(){
  if ! which npx > /dev/null ; then
    echo -e "\nMissing npx. Please download from https://nodejs.org/ja/\n"
    exit 1
  fi
}

# Check swagger-cli exists
function validate_swagger_cli_exists(){
  if ! npx swagger-cli --version > /dev/null ; then
    echo -e "\nMissing swagger-cli. Please run 'npm install'\n"
    exit 1
  fi
}

# Check oapi-codegen exists
function validate_oapi_codegen_exists(){
  if ! which oapi-codegen > /dev/null ; then
    echo -e "\nMissing oapi-codegen. Please download from https://github.com/deepmap/oapi-codegen\n"
    exit 1
  fi
}

# Check goreleaser exists
function validate_goreleaser_exists(){
  if ! which goreleaser > /dev/null ; then
    echo -e "\nMissing goreleaser. Please download from https://goreleaser.com/\n"
    exit 1
  fi
}

# Check .goreleaser.yml exists
function validate_goreleaser_yml_exists(){
  local EXPECT_PATH=$1/.goreleaser.yml
  if [ ! -f "$EXPECT_PATH" ]; then
    echo -e "\nMissing $EXPECT_PATH. Please put it.\n"
    exit 1
  fi
}
