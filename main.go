package main

import (
	"flag"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/jedi/gen"
	"github.com/mh-cbon/jedi/parser"
)

func main() {

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

	base := strings.TrimSuffix(filename, ".go")
	out, err := os.Create(filepath.Join(path, base+"_jedi.go"))
	if err != nil {
		return err
	}
	defer out.Close()

	data := map[string]interface{}{
		"PackageName": filepath.Base(pack),
		"all":         structs,
	}
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
