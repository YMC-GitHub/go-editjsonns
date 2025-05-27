# go-editjsonns

[![Test](https://github.com/ymc-github/go-editjsonns/actions/workflows/test.yml/badge.svg)](https://github.com/ymc-github/go-editjsonns/actions/workflows/test.yml)
[![Release](https://github.com/ymc-github/go-editjsonns/actions/workflows/release.yml/badge.svg)](https://github.com/ymc-github/go-editjsonns/actions/workflows/release.yml)
[![codecov](https://codecov.io/gh/ymc-github/go-editjsonns/branch/main/graph/badge.svg)](https://codecov.io/gh/ymc-github/go-editjsonns)
[![Go Reference](https://pkg.go.dev/badge/github.com/ymc-github/go-editjsonns.svg)](https://pkg.go.dev/github.com/ymc-github/go-editjsonns)

A collection of Go packages for JSON manipulation and processing.

## Packages

### [jsonns](pkg/jsonns)

A package for parsing and manipulating JSON namespace expressions. Provides functionality to work with array-style notation in JSON paths.

Features:
- Parse array-style notation in JSON paths
- Extract array indices and keys
- Standardize namespace strings
- Support for both numeric and string indices

[Read more about jsonns](pkg/jsonns/README.md)

### [jsonctx](pkg/jsonctx)

A package for manipulating JSON data using namespace expressions. Works with nested JSON structures using dot notation and array-style indexing.

Features:
- Get JSON context using namespace notation
- Support for nested object and array access
- Dynamic initialization of objects and arrays
- Flexible namespace separator configuration

[Read more about jsonctx](pkg/jsonctx/README.md)

## Development

### Prerequisites

- Go 1.21 or later
- Docker (optional, for containerized development)

### Building with Docker

```bash
# Build the Docker image
docker build -t go-editjsonns .

# Run tests
docker run --rm go-editjsonns go test ./...

# Development shell
docker run -it --rm -v ${PWD}:/app go-editjsonns sh
```

### Local Development

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./pkg/jsonns
go test ./pkg/jsonctx
```

## Author

- ymc-github <ymc.github@gmail.com>

## License

MIT License

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request 