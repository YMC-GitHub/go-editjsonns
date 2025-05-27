# jsonctx

A Go package for manipulating JSON data using namespace expressions. This package provides functionality to work with nested JSON structures using dot notation and array-style indexing.

## Features

- Get JSON context using namespace notation
- Support for nested object and array access
- Dynamic initialization of objects and arrays
- Flexible namespace separator configuration

## Usage

```go
import "github.com/ymc-github/go-editjsonns/pkg/jsonctx"

// Create a root JSON context
ctx := jsonctx.RootJsonData{}

// Get context for a nested path
result := jsonctx.GetJsonContextInNs("users[0].addresses[1].street", ctx, ".", -1)

// The result contains:
// - Context: The context at the specified path
// - LastNS: The last namespace component
// - Root: The root context with initialized structure

// Initialize a key with a default value
data := map[string]interface{}{}
jsonctx.InitializeKey(data, "config", map[string]interface{}{})
```

## Examples

### Working with Objects

```go
ctx := jsonctx.RootJsonData{}
result := jsonctx.GetJsonContextInNs("user.profile.name", ctx, ".", -1)
// Creates: {"user": {"profile": {}}}
// result.LastNS will be "name"
```

### Working with Arrays

```go
ctx := jsonctx.RootJsonData{}
result := jsonctx.GetJsonContextInNs("users[0].addresses[1]", ctx, ".", -1)
// Creates: {"users": []}
// result.LastNS will be "1"
```

## Integration

This package works in conjunction with the `jsonns` package for namespace standardization and manipulation. 