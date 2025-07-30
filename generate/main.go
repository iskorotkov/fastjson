package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"os/signal"
)

type MyType struct {
	Name string
	Age  int
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get pwd: %w", err)
	}

	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, wd, func(fi fs.FileInfo) bool { return true }, 0)
	if err != nil {
		return fmt.Errorf("parse directory: %w", err)
	}

	for _, pkg := range packages {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}

				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}

					switch v := typeSpec.Type.(type) {
					case *ast.StructType:
						for _, field := range v.Fields.List {
							for _, name := range field.Names {
								fmt.Printf("Found struct field: %s %v\n", name.Name, field.Type)
							}
						}
					case *ast.ArrayType:
					case *ast.MapType:
					}
				}
			}
		}
	}

	return nil
}
