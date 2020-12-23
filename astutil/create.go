package astutil

import (
	"go/ast"
	"go/token"
	"go/types"
)

// CreateNoopOfStatement creates a syntactically safe noop statement out of a given statement.
func CreateNoopOfStatement(pkg *types.Package, info *types.Info, stmt ast.Stmt) ast.Stmt {
	return CreateNoopOfStatements(pkg, info, []ast.Stmt{stmt})
}

// CreateNoopOfStatements creates a syntactically safe noop statement out of a given statement.
func CreateNoopOfStatements(pkg *types.Package, info *types.Info, stmts []ast.Stmt) ast.Stmt {
	var rhs []ast.Expr

	for _, stmt := range stmts {
		for _, identfier := range NoopExprsInStatement(pkg, info, stmt) {
			if removeIdentifierInNooping(identfier) {
				rhs = append(rhs, identfier)
			}
		}
	}

	if len(rhs) == 0 {
		return &ast.EmptyStmt{
			Semicolon: token.NoPos,
		}
	}

	lhs := make([]ast.Expr, len(rhs))
	for i := range rhs {
		lhs[i] = ast.NewIdent("_")
	}

	return &ast.AssignStmt{
		Lhs: lhs,
		Rhs: rhs,
		Tok: token.ASSIGN,
	}
}

func removeIdentifierInNooping(identfier ast.Expr) bool {
	switch node := identfier.(type) {
	case *ast.CompositeLit:
		switch foo := node.Type.(type) {
		case *ast.SelectorExpr:
			if foo.Sel.Obj == nil {
				return false
			}
		}
	}

	return true
}
