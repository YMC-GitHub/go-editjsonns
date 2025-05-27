# jsonns

A Go package for parsing and manipulating JSON namespace expressions. This package provides functionality to work with array-style notation in JSON paths, such as `names[0]` or `users[id]`.

## Features

- Parse array-style notation in JSON paths
- Extract array indices and keys
- Standardize namespace strings
- Support for both numeric and string indices
- Comprehensive test coverage

## Installation

```bash
go get github.com/ymc-github/go-editjsonns/pkg/jsonns
```

## Usage

```go
import "github.com/ymc-github/go-editjsonns/pkg/jsonns"

// Check if a string has array index notation
hasIndex := jsonns.NSHasKeyArrayIndex("names[0]") // true

// Get array index from a string
index := jsonns.NSGetKeyArrayIndex("names[0]") // returns pointer to "0"

// Get key part from array expression
key := jsonns.NSGetKeyArrKey("names[0]") // returns "names"

// Convert to object representation
obj := jsonns.NSKeyarrObjify("names[0]") // returns {Key: "names", Index: "0"}

// Get pure name without brackets
name := jsonns.NSPureName("[0]") // returns "0"

// Find all array expressions
matches := jsonns.NSGetMatch("names[0][1]", nil, "") // returns ["[0]", "[1]"]

// Standardize namespace string
parts := jsonns.NSStd("names[0].users[1]", ".", nil) // returns ["names", "[0]", "users", "[1]"]
```

## Testing

Run the tests with:

```bash
go test
``` 