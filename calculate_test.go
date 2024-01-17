package estimatenecessarytests

import (
	"fmt"
	"go/parser"
	"testing"
)

func Test_NewASTLoader(t *testing.T) {
	loader := NewASTLoader(".", false)
	loader.Load(parser.ParseComments)
	fmt.Println(loader)
}

func TestVCalculator(t *testing.T) {
	loader := NewASTLoader("./testdata", false)
	loader.Load(parser.ParseComments)
	calculator := NewCalculator()

	for _, ast := range loader.Asts {
		calculator.Calculate(ast)
	}

	fmt.Println(calculator)
}
