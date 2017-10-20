//Package main is the jedi cli a golang database generator
package main

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/jedi/gen"
	"github.com/mh-cbon/jedi/parser"
)

// VERSION of the program
var VERSION = "0.0.0"

func showHelp() {
	showVersion()
	fmt.Print(`
A golang database generator to work with dbr (https://github.com/gocraft/dbr)

Usage
	jedi [import packages]

	As a go generator, it looks for environment variables, namely:
		GOFILE: the path to the file containing the //jedi: comments
		GOPACKAGE: the package path related to the GOFILE

`)
}
func showVersion() {
	fmt.Printf(`jedi - %v
`, VERSION)
}

func main() {

	var help bool
	var version bool
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&version, "version", false, "Show version")

	flag.Parse()

	if help {
		showHelp()
		os.Exit(0)
	}
	if version {
		showVersion()
		os.Exit(0)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, arg := range flag.Args() {
		// import arg as directory or package path
		var pack *build.Package
		s, err := os.Stat(arg)
		if err == nil && s.IsDir() {
			pack, err = build.ImportDir(arg, 0)
		}
		if os.IsNotExist(err) {
			err = nil
		}
		if pack == nil && err == nil {
			pack, err = build.Import(arg, wd, 0)
		}
		if err != nil {
			log.Fatalf("%s: %s", arg, err)
		}
		log.Printf("%#v\n", pack)
	}

	file := os.Getenv("GOFILE")
	pack := os.Getenv("GOPACKAGE")
	if file != "" && pack != "" {
		err := processFile(wd, file, pack)
		if err != nil {
			log.Fatalf("%s", err)
		}
		gofmt(wd)
	}
	os.Exit(0)
}

func gofmt(path string) {
	cmd := exec.Command("gofmt", "-s", "-w", path)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("gofmt error: %s %s", err, b)
	}
}

func processFile(path, filename, pack string) error {

	structs, err := parser.Parse(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	if len(structs) == 0 {
		return nil
	}

	data := map[string]interface{}{
		"PackageName": filepath.Base(pack),
		"all":         structs,
	}

	{
		out, err2 := os.Create(filepath.Join(path, "registry_jedi.go"))
		if err2 != nil {
			return err2
		}
		defer out.Close()
		if err2 = gen.Registry.Execute(out, data); err2 != nil {
			return err2
		}
	}

	base := strings.TrimSuffix(filename, ".go")
	out, err := os.Create(filepath.Join(path, base+"_jedi.go"))
	if err != nil {
		return err
	}
	defer out.Close()

	if err = gen.Prolog.Execute(out, data); err != nil {
		return err
	}

	for _, s := range structs {
		data = map[string]interface{}{
			"current": s,
			"all":     structs,
		}
		if err = gen.Struct.Execute(out, data); err != nil {
			return err
		}
	}
	return nil
}
