package main

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
)

// Package is a target package.
type Package struct {
	patterns  []string
	name      string
	typesInfo *types.Info
	syntax    []*ast.File
}

func newPackage(patterns []string) *Package {
	return &Package{patterns: patterns}
}

// Parse a target package.
func (p *Package) Parse() error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes | /* packages.NeedTypesSizes | */
			packages.NeedSyntax | packages.NeedTypesInfo |
			packages.NeedDeps,
		Tests: false,
	}
	pkgs, err := packages.Load(cfg, p.patterns...)
	if err != nil {
		return fmt.Errorf("can't parse package: %w", err)
	}
	if len(pkgs) != 1 {
		return fmt.Errorf("can't parse package: %d packages found", len(pkgs))
	}
	pkg := pkgs[0]
	p.name = pkg.Name
	p.typesInfo = pkg.TypesInfo
	p.syntax = pkg.Syntax
	return nil
}
