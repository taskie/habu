package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:generate habu -type Options,FwvOptions,JcOptions,PityOptions

type Options struct {
	Verbose bool `hb_persistent:"true" hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"verbose mode"`
	Debug   bool `hb_persistent:"true" hb_long:"_" hb_env:"_" hb_usage:"debug mode"`
}

type FwvOptions struct {
	Config           string `hb_persistent:"true" hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"config file"`
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

type JcOptions struct {
	Config   string `hb_persistent:"true" hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"config file"`
	FromType string `hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"convert from [json|toml|yaml|msgpack|dotenv]"`
	ToType   string `hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"convert to [json|toml|yaml|msgpack|dotenv]"`
	Indent   string `hb_short:"I" hb_long:"_" hb_env:"_" hb_usage:"indentation of output"`
}

type PityOptions struct {
	Config string `hb_persistent:"true" hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"config file"`
	Input  string `hb_short:"_" hb_long:"_" hb_value:"pity.txt" hb_env:"_" hb_usage:"pity input file"`
	Output string `hb_short:"_" hb_long:"_" hb_env:"_" hb_usage:"terminal output file"`
}

func unmarshalGlobalOptions() *Options {
	globals := Options{}
	err := viper.Unmarshal(&globals)
	if err != nil {
		log.Fatalf("can't parse global options: %s", err.Error())
	}
	return &globals
}

func fwvRun(cmd *cobra.Command, args []string) {
	globals := unmarshalGlobalOptions()
	log.Printf("%+v", globals)
	options := FwvOptions{}
	err := viper.Unmarshal(&options)
	if err != nil {
		log.Fatalf("can't parse options: %s", err.Error())
	}
	log.Printf("%+v", &options)
}

func jcRun(cmd *cobra.Command, args []string) {
	globals := unmarshalGlobalOptions()
	log.Printf("%+v", globals)
	options := JcOptions{}
	err := viper.Unmarshal(&options)
	if err != nil {
		log.Fatalf("can't parse options: %s", err.Error())
	}
	log.Printf("%+v", &options)
}

func pityRun(cmd *cobra.Command, args []string) {
	globals := unmarshalGlobalOptions()
	log.Printf("%+v", globals)
	options := PityOptions{}
	err := viper.Unmarshal(&options)
	if err != nil {
		log.Fatalf("can't parse options: %s", err.Error())
	}
	log.Printf("%+v", &options)
}

func main() {
	cmd := cobra.Command{
		Use: "subcommands",
	}
	PrepareOptionsFlags(&cmd)
	viper.SetEnvPrefix("SUBCOMMANDS")
	BindOptionsFlags(&cmd, viper.GetViper())

	// You must use PreRun to define flags of a subcommand.
	// See: https://github.com/spf13/viper/issues/233

	fwvCmd := cobra.Command{
		Use: "fwv",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.SetEnvPrefix("FWV")
			BindFwvOptionsFlags(cmd, viper.GetViper())
		},
		Run: fwvRun,
	}
	PrepareFwvOptionsFlags(&fwvCmd)
	cmd.AddCommand(&fwvCmd)

	jcCmd := cobra.Command{
		Use: "jc",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.SetEnvPrefix("JC")
			BindJcOptionsFlags(cmd, viper.GetViper())
		},
		Run: jcRun,
	}
	PrepareJcOptionsFlags(&jcCmd)
	cmd.AddCommand(&jcCmd)

	pityCmd := cobra.Command{
		Use: "pity",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.SetEnvPrefix("PITY")
			BindPityOptionsFlags(cmd, viper.GetViper())
		},
		Run: pityRun,
	}
	PreparePityOptionsFlags(&pityCmd)
	cmd.AddCommand(&pityCmd)

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
