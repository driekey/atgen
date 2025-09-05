# atgen - Go Code Generation Tool

> For Chinese documentation, see [README.zh-CN.md](README.zh-CN.md)

`atgen` is a Go code generation tool that generates field access methods for structs, specifically designed to handle read/write operations for pure data structures.


## Features
- Generates the following method for structs marked with `//go:generate atgen -key=json -type=Obj -output=obj_at.gen.go`:
  - `func (t* Obj) At(key string, visit func(val any) any) error`
- Supports field access via field names or tag mappings
- Properly handles nested structs and anonymous fields
- Ensures type safety by validating that the assigned value matches the field type
- Provides clear error messages when fields are not found or type mismatches occur


## Key Characteristics
- **Type Safety**: Automatically performs type checks to prevent incompatible assignments
- **Error Handling**: Returns detailed error messages for missing fields or failed type assertions
- **Nil Handling Strategy**: Does not update fields when the `visit` function returns `nil` (use dedicated methods for explicit `nil` assignment)
- **Tag Support**: Enables field access via tags such as JSON, YAML, and XML
- **Embedding Support**: Correctly processes nested structs and anonymous fields

> Note: Currently optimized for pure data structs; may not be fully compatible with special types like `func` or `chan`.


## Installation

```bash
go install github.com/driekey/atgen@latest
```


## Usage

1. Add a `go:generate` directive to your struct definition file:
```go
//go:generate atgen -type=YourStruct -key=json -output=your_struct_at.gen.go
```

2. Run code generation:
```bash
go generate ./...
```

3. Use the generated `At` method:
```go
obj := &YourStruct{Name: "test", Age: 30}

// Update field value
err := obj.At("name", func(val any) any {
    return strings.ToUpper(val.(string))
})

// Handle errors
if err != nil {
    log.Printf("update failed: %v", err)
}
```


## Error Handling

The `At` method may return the following errors:
- `field not found: <key>` - Returned when the specified field does not exist
- `type assertion failed for field <key>: expected <type>, got <actual type>` - Returned when a type assertion fails


## Notes
- Use dedicated methods to explicitly set fields to `nil` (since `nil` is used as an indicator for "no update")
- Generated code includes protective comments to prevent overwriting manual edits
- Supports cross-package type references, but ensure dependent packages are properly imported