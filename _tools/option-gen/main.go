package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/c-bata/kube-prompt/internal/optionconv"
	"github.com/k0kubun/pp"
)

var (
	output       string
	pkg          string
	variableName string
)

func convert() error {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	f, err := optionconv.GetOptionsFromHelpText(string(bytes))
	if err != nil {
		return err
	}
	options := optionconv.SplitOptions(f)
	suggests := optionconv.ConvertToSuggestions(options)
	if output == "" {
		_, err = pp.Fprintln(os.Stdout, suggests)
	} else {
		f, err := os.Create(output)
		if err != nil {
			return err
		}
		defer f.Close()

		fmt.Fprintf(f, "// Code generated by 'option-gen'. DO NOT EDIT.\n\n")
		fmt.Fprintf(f, "package %s\n\n", pkg)
		fmt.Fprintln(f, `import (`)
		fmt.Fprintln(f, `prompt "github.com/c-bata/go-prompt"`)
		fmt.Fprintln(f, ")")
		fmt.Fprintln(f, "")
		fmt.Fprintf(f, "var %s = []prompt.Suggest{\n", variableName)
		for _, s := range suggests {
			fmt.Fprintf(f, "%s,\n", s.String())
		}
		fmt.Fprintln(f, "}")
	}
	return err
}

func main() {
	flag.StringVar(&output, "o", "", "output file. print stdout if empty")
	flag.StringVar(&pkg, "pkg", "kube", "package name")
	flag.StringVar(&variableName, "var", "flagXXX", "variable name")
	flag.Parse()

	if err := convert(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
