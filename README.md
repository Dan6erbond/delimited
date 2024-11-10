# 〰️ delimited

A Go library to marshal/unmarshal delimited strings.

## Introduction

〰️ delimited is a flexible Go library for marshaling and unmarshaling structs into a compact string format using a customizable delimiter. It’s ideal for applications that benefit from simplified data representations, avoiding the overhead of formats like JSON or YAML.

This can, for example, be useful for GraphQL APIs providing a cursor to clients used to paginate through large datasets.

## Features

- [x] Marshalling/unmarshalling delimited strings with custom delimiter using `DecoderWithDelimiter()` and `EncoderWithDelimiter()` options.
- [x] Ignoring struct fields with `delimited:"ignore"` tag.
- [x] Parsing custom data types using `json` package.

## Example

See the test files for examples on how to use 〰️ delimited.

A specific cursor example can be found at [`cursor_test.go`](/cursor_test.go).
