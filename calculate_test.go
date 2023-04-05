package estimatenecessarytests

import (
	"fmt"
	"testing"
)

func Test_NewASTLoader(t *testing.T) {
	loader := NewASTLoader(".", false)
	loader.Load()
	fmt.Println(loader)
}

func TestVCalculator(t *testing.T) {
	loader := NewASTLoader("./testdata", false)
	loader.Load()
	calculator := NewCalculator()

	for _, ast := range loader.asts {
		calculator.Calculate(ast)
	}

	fmt.Println(calculator)
}
