package main

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:generate habu -type Options

type Options struct {
	Long       string `hb_long:"some-long"`
	Short      string `hb_long:"some-short" hb_short:"S"`
	Usage      string `hb_long:"_" hb_short:"_" hb_usage:"Some usage"`
	Env        string `hb_long:"_" hb_short:"_" hb_env:"SOME_ENV"`
	Persistent string `hb_persistent:"true" hb_long:"_" hb_env:"_"`

	String   string        `hb_long:"_" hb_value:"hello"`
	Int      int           `hb_long:"_" hb_value:"42"`
	Bool     bool          `hb_long:"_" hb_value:"true"`
	Duration time.Duration `hb_long:"_" hb_value:"1h"`
}

func run(cmd *cobra.Command, args []string) {
	options := Options{}
	err := viper.Unmarshal(&options)
	if err != nil {
		log.Fatalf("can't parse options: %s", err.Error())
	}
	log.Printf("%+v", &options)
}

func main() {
	cmd := cobra.Command{
		Use: "features",
		Run: run,
	}
	PrepareOptionsFlags(&cmd)
	viper.SetEnvPrefix("FEATURES")
	BindOptionsFlags(&cmd, viper.GetViper())
	cmd.Execute()
}
