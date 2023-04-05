package main

import (
	"flag"
	"fmt"

	estimatenecessarytests "github.com/lkeix/estimate-necessary-tests"
)

func main() {
	var path string
	flag.StringVar(&path, "path", ".", "specify directory estimate number of necessary tests")
	flag.Parse()

	loader := estimatenecessarytests.NewASTLoader(path, false)
	loader.Load()

	calculator := estimatenecessarytests.NewCalculator()
	for _, ast := range loader.Asts {
		calculator.Calculate(ast)
	}

	fmt.Printf("output format ${package}.${object}.${function}.\n${object} is optional.\n")
	for key, res := range calculator.Result {
		fmt.Printf("%s needs %d tests\n", key, res)
	}
}
