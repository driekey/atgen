# atgen - Go 代码生成工具

`atgen` 是一个为结构体生成读写访问方法的 Go 代码生成工具。

## 功能
- 为带有 `//go:generate atgen -key=json -type=Obj -output=obj_at.gen.go` 标记的结构体生成方法：
  - `func (t* Obj) At(key string, visit func(val any) any)`
- 支持通过字段名或标签映射访问字段
- 正确处理嵌套结构体和匿名字段
- 自动处理指针类型，支持基础类型到指针类型的自动转换
- 类型安全检查，确保写入值的类型与字段类型匹配

> 目前是用来给定义纯数据结构体而使用的, 对于存在的特殊类型如 func, chan 可能不适配

## 安装

```bash
go install github.com/driekey/atgen@latest
```