package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
	"strconv"
	"time"
)

// Collector is a collector of configurations for pflag.FlagSet .
type Collector struct {
	pkg  *Package
	errs []error
}

func newCollector(pkg *Package) *Collector {
	return &Collector{pkg: pkg}
}

// Collect configurations.
func (c *Collector) Collect(typeNames []string) ([]*FlagSet, []error) {
	c.errs = make([]error, 0)
	flagSets := make([]*FlagSet, 0)
	for _, typeName := range typeNames {
		flagSet := c.collect(typeName)
		if flagSet == nil {
			continue
		}
		flagSets = append(flagSets, flagSet)
	}
	return flagSets, c.errs
}

func (c *Collector) Errorf(format string, a ...interface{}) {
	c.errs = append(c.errs, fmt.Errorf(format, a...))
}

func (c *Collector) collect(typeName string) *FlagSet {
	for _, file := range c.pkg.syntax {
		if file != nil {
			fc := FlagSetCollector{
				parent: c,
				file:   file,
				flagSet: &FlagSet{
					typeName: typeName,
					flags:    make([]*Flag, 0),
				},
			}
			ast.Inspect(file, fc.collect)
			if len(fc.flagSet.flags) != 0 {
				return fc.flagSet
			}
		}
	}
	return nil
}

type FlagSetCollector struct {
	parent  *Collector
	file    *ast.File
	flagSet *FlagSet
}

func (c *FlagSetCollector) Errorf(format string, a ...interface{}) {
	c.parent.Errorf(format, a...)
}

func (c *FlagSetCollector) collect(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok {
		return true
	}
	for _, spec := range decl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		ident := typeSpec.Name
		if ident.Name != c.flagSet.typeName {
			continue
		}
		def := c.parent.pkg.typesInfo.ObjectOf(ident)
		if def == nil {
			continue
		}
		typeName, ok := def.(*types.TypeName)
		if !ok {
			continue
		}
		named, ok := typeName.Type().(*types.Named)
		if !ok {
			continue
		}
		st, ok := named.Underlying().(*types.Struct)
		if !ok {
			continue
		}
		n := st.NumFields()
		for i := 0; i < n; i++ {
			v := st.Field(i)
			if v.Anonymous() || !v.IsField() {
				continue
			}
			tag := reflect.StructTag(st.Tag(i))
			name := tag.Get("hb_long")
			env := tag.Get("hb_env")
			if name == "" && env == "" {
				continue
			}
			f := c.collectFlag(name, env, v, tag)
			if f == nil {
				continue
			}
			c.flagSet.flags = append(c.flagSet.flags, f)
		}
	}
	return false
}

func (c *FlagSetCollector) collectFlag(name, env string, v *types.Var, tag reflect.StructTag) *Flag {
	fieldName := v.Name()
	f := &Flag{
		fieldName: fieldName,
		name:      name,
		shorthand: tag.Get("hb_short"),
		usage:     tag.Get("hb_usage"),
		parse:     tag.Get("hb_parse"),
		viper:     tag.Get("hb_viper"),
		env:       env,
	}
	var err error
	persistent := tag.Get("hb_persistent")
	if persistent != "" {
		f.persistent, err = strconv.ParseBool(persistent)
		if err != nil {
			c.Errorf("invalid hb_persistent value: %s (%s.%s)", persistent, c.flagSet.typeName, f.fieldName)
		}
	}
	flagValue := tag.Get("hb_value")
	switch ty := v.Type().(type) {
	case *types.Basic:
		switch ty.Kind() {
		case types.Bool:
			f.fieldType = Bool
			if flagValue != "" {
				f.value, err = strconv.ParseBool(flagValue)
				if err != nil {
					c.Errorf("invalid %s value: %v (%s.%s)", f.fieldType, err, c.flagSet.typeName, f.fieldName)
				}
			} else {
				f.value = false
			}
		case types.Int:
			f.fieldType = Int
			if flagValue != "" {
				f.value, err = strconv.Atoi(flagValue)
				if err != nil {
					c.Errorf("invalid %s value: %v (%s.%s)", f.fieldType, err, c.flagSet.typeName, f.fieldName)
				}
			} else {
				f.value = 0
			}
		case types.String:
			f.fieldType = String
			f.value = flagValue
		}
	case *types.Named:
		obj := ty.Obj()
		if obj.Pkg().Path() == "time" && obj.Id() == "Duration" {
			f.fieldType = Duration
			if flagValue != "" {
				f.value, err = time.ParseDuration(flagValue)
				if err != nil {
					c.Errorf("invalid %s value: %v (%s.%s)", f.fieldType, err, c.flagSet.typeName, f.fieldName)
				}
			} else {
				f.value = time.Duration(0)
			}
		}
	}
	if f.fieldType == None {
		return nil
	}
	return f
}
