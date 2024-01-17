package estimatenecessarytests

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
)

type ASTLoader struct {
	files []string
	Asts  map[string]*ast.File
}

func NewASTLoader(path string, includeTest bool) *ASTLoader {
	re := regexp.MustCompile(`\.go$`)
	testRe := regexp.MustCompile(`\_test.go$`)
	files := []string{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && re.MatchString(info.Name()) {
			if !includeTest && testRe.MatchString(info.Name()) {
				return nil
			}
			files = append(files, path)
		}

		return nil
	})

	return &ASTLoader{
		files: files,
		Asts:  make(map[string]*ast.File),
	}
}

func (a *ASTLoader) Load(parseMode parser.Mode) error {
	for _, f := range a.files {
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, f, nil, parseMode)
		if err != nil {
			return err
		}

		a.Asts[f] = node
	}
	return nil
}

type Calculator struct {
	Result map[string]int64
}

func NewCalculator() *Calculator {
	return &Calculator{
		Result: make(map[string]int64),
	}
}

func (c *Calculator) Calculate(node *ast.File) {
	for _, decl := range node.Decls {
		if f, ok := decl.(*ast.FuncDecl); ok {
			needTests := parseFunc(f)
			funcName := f.Name.Name
			pkgName := node.Name.Name
			structName := extractStructName(f)
			key := ""
			if structName != "" {
				key = fmt.Sprintf("%s.%s.%s", pkgName, structName, funcName)
			} else {
				key = fmt.Sprintf("%s.%s", pkgName, funcName)
			}
			c.Result[key] = needTests
		}
	}
}

func extractStructName(f *ast.FuncDecl) string {
	if f.Recv != nil {
		if len(f.Recv.List) > 0 {
			if t, ok := f.Recv.List[0].Type.(*ast.StarExpr); ok {
				if ident, ok := t.X.(*ast.Ident); ok {
					return ident.Name
				}
			}
		}
	}
	return ""
}

func parseFunc(f *ast.FuncDecl) int64 {
	return calculate(f.Body.List, 1)
}

func calculate(stmts []ast.Stmt, cnt int64) int64 {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *ast.ForStmt:
			cnt = calculate(s.Body.List, cnt)
		case *ast.IfStmt:
			if s.Else != nil {
				cnt++
				switch el := s.Else.(type) {
				case *ast.IfStmt:
					cnt = calculate(el.Body.List, cnt)
				}
			}

			if s.Cond != nil {
				cnt++
			}
			cnt = calculate(s.Body.List, cnt)
		case *ast.SwitchStmt:
			cnt = calculate(s.Body.List, cnt)
		case *ast.CaseClause:
			cnt = calculate(s.Body, cnt+1)
		}
	}
	return cnt
}
