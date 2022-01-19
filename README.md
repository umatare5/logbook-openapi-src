# logbook-openapi-src

logbook-openapi-src is a source generates libraries for [Logbook API](https://github.com/umatare5/logbook-api).

This code generates following two modules;

- [logbook-go](https://github.com/umatare5/logbook-go)

  Go Client to use Logbook API.

- [logbook-api-framework-impl](https://github.com/umatare5/logbook-api-framework-impl)

  Implementation for framework layer in Logbook API.

## Development

### Setup

- Install `swagger-cli` for bundle divided OpenAPI Specifications.

  ```sh
  npm install
  ```

- Install `oapi-codegen` for build the client.

  ```sh
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
  ```

### Build

```sh
make build
```

### Release

```sh
git bump
make release
```
