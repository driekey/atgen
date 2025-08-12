package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	typeFlag := flag.String("type", "", "Tag type (e.g. json, yaml)")
	outputFlag := flag.String("output", "", "Output file name")
	recursiveFlag := flag.Bool("recursive", false, "Process directories recursively")
	flag.Parse()

	// 获取文件/目录列表
	paths := flag.Args()
	if len(paths) == 0 {
		if gofile := os.Getenv("GOFILE"); gofile != "" {
			paths = []string{gofile}
		} else {
			fmt.Println("Error: No input files specified")
			os.Exit(1)
		}
	}

	// 收集所有Go文件（递归或非递归）
	var files []string
	for _, path := range paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Printf("Error resolving path %q: %v\n", path, err)
			os.Exit(1)
		}

		err = filepath.Walk(absPath, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// 跳过目录（除非是根目录）
			if info.IsDir() {
				if p != absPath && !*recursiveFlag {
					return filepath.SkipDir // 非递归模式跳过子目录
				}
				return nil
			}

			// 只处理.go文件且排除测试文件
			if strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "_test.go") {
				files = append(files, p)
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error walking path %q: %v\n", absPath, err)
			os.Exit(1)
		}
	}

	// 检查是否找到文件
	if len(files) == 0 {
		fmt.Println("Error: No valid .go files found")
		os.Exit(1)
	}

	config := Config{
		TagType: *typeFlag,
		Output:  *outputFlag,
	}

	gen := NewGenerator(config)
	if err := gen.ProcessFiles(files); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
