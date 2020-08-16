package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// See also: https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go

var (
	typeNames = flag.String("type", "", "comma-separated list of type names; must be set")
	output    = flag.String("output", "", "output file name; default srcdir/<type>_string.go")
	viper     = flag.String("viper", "", "viper mode")
	longCase  = flag.String("longCase", "kebab", "auto case conversion for long options")
	shortCase = flag.String("shortCase", "lower", "auto case conversion for short options")
	envCase   = flag.String("envCase", "", "auto case conversion for environment variables")
	envPrefix = flag.String("envPrefix", "", "auto prefix for environment variables")
	debug     = flag.Bool("debug", false, "show debug output")
)

var spaces = regexp.MustCompile("\\s+")

func main0() error {
	log.SetFlags(0)
	log.SetPrefix("habu: ")
	flag.Parse()
	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	types := strings.Split(*typeNames, ",")

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	var dir string
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		dir = filepath.Dir(args[0])
	}

	pkg := newPackage(args)
	err := pkg.Parse()
	if err != nil {
		log.Fatal(err)
	}
	c := newCollector(pkg)
	flagSets, errs := c.Collect(types)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err.Error())
		}
		panic(errs[0])
	}
	g := newGenerator(pkg.name, flagSets)
	src := g.Generate()

	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("%s_habu.go", types[0])
		outputName = filepath.Join(dir, strings.ToLower(baseName))
	}
	return ioutil.WriteFile(outputName, src, 0644)
}

func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

func main() {
	err := main0()
	if err != nil {
		panic(err)
	}
}
