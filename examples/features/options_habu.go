// Code generated by "habu -type Options -viper true"; DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func PrepareOptionsFlags(c *cobra.Command, v *viper.Viper) {
	fs := c.Flags()
	ps := c.PersistentFlags()
	fs.String("some-long", "", "")
	v.BindPFlag("Long", fs.Lookup("some-long"))
	fs.StringP("some-short", "S", "", "")
	v.BindPFlag("Short", fs.Lookup("some-short"))
	fs.StringP("usage", "u", "", "Some usage")
	v.BindPFlag("Usage", fs.Lookup("usage"))
	fs.StringP("env", "e", "", "")
	v.BindPFlag("Env", fs.Lookup("env"))
	v.BindEnv("Env", "SOME_ENV")
	ps.String("persistent", "", "")
	v.BindPFlag("Persistent", ps.Lookup("persistent"))
	v.BindEnv("Persistent")
	fs.String("string", "hello", "")
	v.BindPFlag("String", fs.Lookup("string"))
	fs.Int("int", 42, "")
	v.BindPFlag("Int", fs.Lookup("int"))
	fs.Bool("bool", true, "")
	v.BindPFlag("Bool", fs.Lookup("bool"))
	fs.Duration("duration", 3600000000000, "")
	v.BindPFlag("Duration", fs.Lookup("duration"))
}