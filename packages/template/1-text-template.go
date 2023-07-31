package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	sprig "github.com/Masterminds/sprig/v3"
)

func templatingOne() error {
	// "." or dot acts like this of other language (my opinion :p)
	tmplStr := "My name is {{.}}.\n"
	name := "Ishtiaq"
	tmpl, err := template.New("test").Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		return err
	}
	return nil
}

// fields must be exported to use template on this struct
type Invertory struct {
	Material string
	Count    int
}

func templatingTwo() error {
	sweaters := Invertory{"wool", 17}
	tmplStr := "{{.Count}} items are made of {{.Material}}.\n"
	tmpl, err := template.New("test").Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		return err
	}
	return nil
}

func templatingThree() error {
	// conditional pipeline (we can use if, if...else, if...else if)
	tmplStr := "My name is {{- if .}} Ishtiaq {{- end}}.\n"
	print := "Yes" // zero value will not print the name (print := "")
	tmpl, err := template.New("test").Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, print)
	if err != nil {
		return err
	}
	return nil
}

func templatingFour() error {
	// range over pipeline
	tmplStr := "Count 1 to 5: {{range .}}\n{{.}}{{end}}\n"
	nums := []int{1, 2, 3, 4, 5}
	tmpl, err := template.New("test").Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, nums)
	if err != nil {
		return err
	}
	return nil
}

func templatingFive() error {
	// variables
	tmplStr := "Count 1 to 5: {{range $index, $element := .}}\nslice[{{$index}}] = {{$element}}{{end}}\n"
	nums := []int{1, 2, 3, 4, 5}
	tmpl, err := template.New("test").Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, nums)
	if err != nil {
		return err
	}
	return nil
}

func templatingSix() error {
	// Functions (default)
	tmplStr := "Count 1 to 5 (except index 3): {{range $index, $element := .}}{{if eq $index 3}}{{continue}}{{else}}\nslice[{{$index}}] = {{$element}}{{end}}{{end}}\n"
	nums := []int{1, 2, 3, 4, 5}
	tmpl, err := template.New("test").Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, nums)
	if err != nil {
		return err
	}
	return nil
}

func templatingSeven() error {
	// Functions (custom)
	tmplStr := "Print 5 names in uppercase: {{range $index, $element := .}}\nslice[{{$index}}] = {{upperCase $element}}{{end}}\n"
	names := []string{"xyz", "abc", "mno", "rst", "pno"}
	funcMap := map[string]interface{}{
		"upperCase": strings.ToUpper,
	}
	tmpl, err := template.New("test").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, names)
	if err != nil {
		return err
	}
	return nil
}

func templatingEight() error {
	// Functions (sprig)
	// Useful template functions for Go templates.
	// https://github.com/Masterminds/sprig
	// http://masterminds.github.io/sprig/
	tmplStr := "Print 5 names in uppercase: {{range $index, $element := .}}\nslice[{{$index}}] = {{upper $element}}{{end}}\n"
	names := []string{"xyz", "abc", "mno", "rst", "pno"}
	tmpl, err := template.New("test").Funcs(sprig.FuncMap()).Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, names)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Text Templating!!!")

	err := templatingOne()
	if err != nil {
		panic(err)
	}

	err = templatingTwo()
	if err != nil {
		panic(err)
	}

	err = templatingThree()
	if err != nil {
		panic(err)
	}

	err = templatingFour()
	if err != nil {
		panic(err)
	}

	err = templatingFive()
	if err != nil {
		panic(err)
	}

	err = templatingSix()
	if err != nil {
		panic(err)
	}

	err = templatingSeven()
	if err != nil {
		panic(err)
	}

	err = templatingEight()
	if err != nil {
		panic(err)
	}

	err = testing()
	if err != nil {
		panic(err)
	}
}

func testing() error {
	tmplStr := "{{- $msg := dict \"name\" \"Ishtiaq\" \"age\" \"25\" -}} { \"blocks\": [{\"type\": \"intro\", \"text\": {{ toJson $msg }}}]}"
	name := "hello"
	tmpl, err := template.New("test").Funcs(sprig.TxtFuncMap()).Parse(tmplStr)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		return err
	}
	return nil
}
