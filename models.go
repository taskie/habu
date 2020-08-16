package main

type FlagSet struct {
	typeName string
	flags    []*Flag
}

type Flag struct {
	fieldName string
	fieldType FlagType

	name      string
	shorthand string
	value     interface{}
	usage     string

	parse      string
	persistent bool
	viper      string
	env        string
}

//go:generate stringer -type=FlagType

type FlagType int

const (
	None FlagType = iota

	Int
	Bool
	String

	Duration
)
