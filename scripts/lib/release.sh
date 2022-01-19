# Release client
function release_library(){
  local GO_RELEASE_DIR=$1
  local GIT_LATEST_TAG=$2

  cd ${GO_RELEASE_DIR}
  git add .
  git commit -m "Bump library to ${GO_RELEASE_DIR}"
  echo ${GIT_LATEST_TAG}
  git tag ${GIT_LATEST_TAG}
  git push origin main

  goreleaser release --rm-dist
  cd - > /dev/null
}
