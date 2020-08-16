package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:generate habu -type Options

type Options struct {
	Config           string `hb_short:"_" hb_long:"_" hb_env:"_" hb_persistent:"true" hb_usage:"config file"`
	FromType         string `hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"convert from [fwv|csv]"`
	ToType           string `hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"convert to [fwv|csv]"`
	NoWidth          bool   `hb_short:"W" hb_long:"_" hb_env:"_" hb_usage:"NOT use char width"`
	EaaHalfWidth     bool   `hb_short:"E" hb_long:"_" hb_env:"_" hb_usage:"treat East Asian Ambiguous as half width"`
	ShowColumnRanges bool   `hb_short:"r" hb_long:"_" hb_env:"_" hb_usage:"show column ranges"`
	NoTrim           bool   `hb_short:"T" hb_long:"_" hb_env:"_" hb_usage:"NOT trim whitespaces"`
	Color            bool   `hb_short:"C" hb_long:"_" hb_env:"_" hb_usage:"colorize output"`
	NoColor          bool   `hb_short:"M" hb_long:"_" hb_env:"NO_COLOR" hb_usage:"NOT colorize output (monochrome)"`
	Whitespaces      string `hb_short:"s" hb_long:"_" hb_env:"_" hb_value:" " hb_usage:"characters treated as whitespace"`
	Delimiter        string `hb_short:"d" hb_long:"_" hb_env:"_" hb_value:" " hb_usage:"delimiter used for FWV output"`
}

func run(cmd *cobra.Command, args []string) {
	options := Options{}
	err := viper.Unmarshal(&options)
	if err != nil {
		log.Fatal("can't parse options: %s", err.Error())
	}
	log.Printf("%+v", options)
}

func main() {
	cmd := cobra.Command{
		Use: "fwv",
		Run: run,
	}
	viper.SetEnvPrefix("FWV")
	PrepareOptionsFlags(&cmd, viper.GetViper())
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
