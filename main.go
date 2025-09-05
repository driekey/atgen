package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

var verbose = flag.Bool("v", false, "enable verbose logging")

func main() {
	typeNames := flag.String("type", "", "comma-separated list of struct type names (optional)")
	keyType := flag.String("key", "field", "key type (field, json, yaml)")
	outputName := flag.String("output", "", "output file name (default: stdout)")
	flag.Parse()

	// Get current file information
	goFile := os.Getenv("GOFILE")
	if goFile == "" {
		log.Fatal("must run in go generate context")
	}

	// Get package path
	pkgPath, err := getPackagePath()
	if err != nil {
		log.Fatalf("failed to get package path: %v", err)
	}

	// If no type specified, parse all structs in current file
	var types []string
	if *typeNames == "" {
		types, err = findStructsInFile(goFile)
		if err != nil {
			log.Fatalf("failed to parse structs in file: %v", err)
		}
	} else {
		types = strings.Split(*typeNames, ",")
		for i := range types {
			types[i] = strings.TrimSpace(types[i])
		}
	}

	if len(types) == 0 {
		logVerbose("no structs found for generation")
		return
	}

	// Generate code for each type
	for _, typeName := range types {
		g := NewGenerator(typeName, *keyType, pkgPath)
		outputFile := *outputName
		if outputFile == "" {
			outputFile = fmt.Sprintf("%s_at.gen.go", toSnakeCase(typeName))
		}

		if err := g.Generate(outputFile); err != nil {
			log.Fatalf("failed to generate code for type %s: %v", typeName, err)
		}

		logVerbose("generated %s for type %s", outputFile, typeName)
	}
}

// logVerbose prints formatted log message only in verbose mode
func logVerbose(format string, v ...any) {
	if *verbose {
		log.Printf("[atgen][verbose] %s", fmt.Sprintf(format, v...))
	}
}

// Get package path
func getPackagePath() (string, error) {
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}

	cfg := &packages.Config{
		Dir:  wd,
		Mode: packages.NeedModule,
		Env:  os.Environ(),
	}

	// load packages
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		return "", fmt.Errorf("failed to load packages: %w", err)
	}

	if len(pkgs) == 0 || pkgs[0].Module == nil {
		return "", fmt.Errorf("no module found (missing go.mod?)")
	}

	mod := pkgs[0].Module
	logVerbose("found module: %s (path: %s, dir: %s)", mod.Path, mod.Dir, wd)

	relPath, err := filepath.Rel(mod.Dir, wd)
	if err != nil {
		return "", fmt.Errorf("failed to calculate relative path: %w", err)
	}

	if relPath == "." {
		return mod.Path, nil
	}
	return filepath.Join(mod.Path, relPath), nil
}
