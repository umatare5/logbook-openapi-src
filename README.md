# logbook-go

logbook-go is Go library to use [Logbook API](https://github.com/umatare5/logbook-api).

## Usage

```sh
go get github.com/umatare5/logbook-go
```

## Features

- Get profile
- Fetch divelogs

## Development

The code is generated from OpenAPI Specification.

### Setup

- Install `swagger-cli` for bundle divided OpenAPI Specifications.

  ```sh
  npm install
  ```

- Install `oapi-codegen` for build the client.

  ```sh
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
  ```

## Operation

### Build

```sh
scripts/build.sh
```

### Release

```sh
scripts/release.sh
```
