# atgen - Go 代码生成工具

`atgen` 是一个为结构体生成 `At` 方法的 Go 代码生成工具。

## 功能
- 为带有 `//go:generate` 标记的结构体生成 `func (t* Obj)At(key string) (val any, ok bool)` 方法
  用于通过 key 访问内部字段
- 支持通过字段名或标签映射访问字段
- 正确处理嵌套结构体和匿名字段

## 安装

```bash
go install github.com/driekey/atgen@latest
```