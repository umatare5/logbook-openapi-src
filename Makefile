.PHONY: build
build:
	./scripts/build.sh

.PHONY: release
release: build
	./scripts/release.sh
