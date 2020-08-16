// Code generated by "habu -type Options,FwvOptions,JcOptions,PityOptions"; DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func PrepareOptionsFlags(c *cobra.Command) {
	ps := c.PersistentFlags()
	ps.BoolP("verbose", "v", false, "verbose mode")
	ps.Bool("debug", false, "debug mode")
}

func BindOptionsFlags(c *cobra.Command, v *viper.Viper) {
	ps := c.PersistentFlags()
	v.BindPFlag("Verbose", ps.Lookup("verbose"))
	v.BindEnv("Verbose")
	v.BindPFlag("Debug", ps.Lookup("debug"))
	v.BindEnv("Debug")
}

func PrepareFwvOptionsFlags(c *cobra.Command) {
	fs := c.Flags()
	ps := c.PersistentFlags()
	ps.StringP("config", "c", "", "config file")
	fs.StringP("from-type", "f", "", "convert from [fwv|csv]")
	fs.StringP("to-type", "t", "", "convert to [fwv|csv]")
	fs.BoolP("no-width", "W", false, "NOT use char width")
	fs.BoolP("eaa-half-width", "E", false, "treat East Asian Ambiguous as half width")
	fs.BoolP("show-column-ranges", "r", false, "show column ranges")
	fs.BoolP("no-trim", "T", false, "NOT trim whitespaces")
	fs.BoolP("color", "C", false, "colorize output")
	fs.BoolP("no-color", "M", false, "NOT colorize output (monochrome)")
	fs.StringP("whitespaces", "s", " ", "characters treated as whitespace")
	fs.StringP("delimiter", "d", " ", "delimiter used for FWV output")
}

func BindFwvOptionsFlags(c *cobra.Command, v *viper.Viper) {
	fs := c.Flags()
	ps := c.PersistentFlags()
	v.BindPFlag("Config", ps.Lookup("config"))
	v.BindEnv("Config")
	v.BindPFlag("FromType", fs.Lookup("from-type"))
	v.BindEnv("FromType")
	v.BindPFlag("ToType", fs.Lookup("to-type"))
	v.BindEnv("ToType")
	v.BindPFlag("NoWidth", fs.Lookup("no-width"))
	v.BindEnv("NoWidth")
	v.BindPFlag("EaaHalfWidth", fs.Lookup("eaa-half-width"))
	v.BindEnv("EaaHalfWidth")
	v.BindPFlag("ShowColumnRanges", fs.Lookup("show-column-ranges"))
	v.BindEnv("ShowColumnRanges")
	v.BindPFlag("NoTrim", fs.Lookup("no-trim"))
	v.BindEnv("NoTrim")
	v.BindPFlag("Color", fs.Lookup("color"))
	v.BindEnv("Color")
	v.BindPFlag("NoColor", fs.Lookup("no-color"))
	v.BindEnv("NoColor", "NO_COLOR")
	v.BindPFlag("Whitespaces", fs.Lookup("whitespaces"))
	v.BindEnv("Whitespaces")
	v.BindPFlag("Delimiter", fs.Lookup("delimiter"))
	v.BindEnv("Delimiter")
}

func PrepareJcOptionsFlags(c *cobra.Command) {
	fs := c.Flags()
	ps := c.PersistentFlags()
	ps.StringP("config", "c", "", "config file")
	fs.StringP("from-type", "f", "", "convert from [json|toml|yaml|msgpack|dotenv]")
	fs.StringP("to-type", "t", "", "convert to [json|toml|yaml|msgpack|dotenv]")
	fs.StringP("indent", "I", "", "indentation of output")
}

func BindJcOptionsFlags(c *cobra.Command, v *viper.Viper) {
	fs := c.Flags()
	ps := c.PersistentFlags()
	v.BindPFlag("Config", ps.Lookup("config"))
	v.BindEnv("Config")
	v.BindPFlag("FromType", fs.Lookup("from-type"))
	v.BindEnv("FromType")
	v.BindPFlag("ToType", fs.Lookup("to-type"))
	v.BindEnv("ToType")
	v.BindPFlag("Indent", fs.Lookup("indent"))
	v.BindEnv("Indent")
}

func PreparePityOptionsFlags(c *cobra.Command) {
	fs := c.Flags()
	ps := c.PersistentFlags()
	ps.StringP("config", "c", "", "config file")
	fs.StringP("input", "i", "pity.txt", "pity input file")
	fs.StringP("output", "o", "", "terminal output file")
}

func BindPityOptionsFlags(c *cobra.Command, v *viper.Viper) {
	fs := c.Flags()
	ps := c.PersistentFlags()
	v.BindPFlag("Config", ps.Lookup("config"))
	v.BindEnv("Config")
	v.BindPFlag("Input", fs.Lookup("input"))
	v.BindEnv("Input")
	v.BindPFlag("Output", fs.Lookup("output"))
	v.BindEnv("Output")
}
