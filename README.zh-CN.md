# atgen - Go 代码生成工具

`atgen` 是一个为结构体生成字段访问方法的 Go 代码生成工具，特别适合处理纯数据结构的读写操作。

## 功能
- 为带有 `//go:generate atgen -key=json -type=Obj -output=obj_at.gen.go` 标记的结构体生成方法：
  - `func (t* Obj) At(key string, visit func(val any) any) error`
- 支持通过字段名或标签映射访问字段
- 正确处理嵌套结构体和匿名字段
- 类型安全检查，确保写入值的类型与字段类型匹配
- 错误处理：当字段不存在或类型不匹配时返回明确的错误信息

## 特性
- **类型安全**：自动进行类型检查，防止不兼容的赋值操作
- **错误处理**：当字段不存在或类型断言失败时返回详细的错误信息
- **nil处理策略**：visit函数返回nil时不进行赋值操作（需要使用其他方法显式设置nil值）
- **标签支持**：支持通过JSON、YAML、XML等标签名访问字段
- **嵌入支持**：正确处理嵌套结构体和匿名字段

> 注意：目前主要用于处理纯数据结构体，对于特殊类型如func、chan可能不完全适配

## 安装

```bash
go install github.com/driekey/atgen@latest
```

## 使用方法

1. 在您的结构体定义文件中添加go generate指令：
```go
//go:generate atgen -type=YourStruct -key=json -output=your_struct_at.gen.go
```

2. 运行代码生成：
```bash
go generate ./...
```

3. 使用生成的At方法：
```go
obj := &YourStruct{Name: "test", Age: 30}

// 更新字段值
err := obj.At("name", func(val any) any {
    return strings.ToUpper(val.(string))
})

// 处理错误
if err != nil {
    log.Printf("update failed: %v", err)
}
```

## 错误处理

At方法可能返回以下错误：
- `field not found: <key>` - 当指定的字段不存在时
- `type assertion failed for field <key>: expected <type>, got <actual type>` - 当类型断言失败时

## 注意事项

- 如果需要显式设置字段为nil，需要使用其他方法，因为nil被用作"不更新"的指示符
- 生成的代码包含保护性注释，防止手动编辑被覆盖
- 支持跨包类型引用，但需要确保相关包已正确导入