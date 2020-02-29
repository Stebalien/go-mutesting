package structinit

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/AntonStoeckl/go-mutesting/mutator"
)

func init() {
	mutator.Register("structinit/remove", MutatorStructinitRemove)
}

// MutatorStructinitRemove implements a mutator to remove fields from struct initialisation.
func MutatorStructinitRemove(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	var mutations []mutator.Mutation

	if blockStmt, ok := node.(*ast.BlockStmt); ok {
		for blockStmtExprIdx, blockStmtExpr := range blockStmt.List {
			if assignStmt, ok := blockStmtExpr.(*ast.AssignStmt); ok {
				if compositeLit, ok := assignStmt.Rhs[0].(*ast.CompositeLit); ok {
					ident, ok := compositeLit.Type.(*ast.Ident)
					if !ok || ident.Obj == nil {
						continue
					}

					typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec)
					if !ok {
						continue
					}

					if _, ok := typeSpec.Type.(*ast.StructType); !ok {
						continue
					}

					for compositeLitElementIdx := range compositeLit.Elts {
						originalCompositeLitElements := compositeLit.Elts
						var newCompositeLitElements []ast.Expr
						var assignToVoid *ast.AssignStmt

						for i := 0; i < len(compositeLit.Elts); i++ {
							// remove each compositeLitElement once (by not adding it to Elts)
							if compositeLitElementIdx != i {
								newCompositeLitElements = append(newCompositeLitElements, compositeLit.Elts[i])
							}

							if compositeLitElementIdx == i {
								if value, ok := compositeLit.Elts[i].(*ast.KeyValueExpr).Value.(*ast.Ident); ok {
									assignToVoid = &ast.AssignStmt{
										Lhs: []ast.Expr{ast.NewIdent("_")},
										Rhs: []ast.Expr{value},
										Tok: token.ASSIGN,
									}
								}
							}
						}

						insertIdx := blockStmtExprIdx + 1

						mutations = append(mutations, mutator.Mutation{
							Change: func() {
								compositeLit.Elts = newCompositeLitElements

								if assignToVoid != nil {
									blockStmt.List = append(blockStmt.List, nil)
									copy(blockStmt.List[insertIdx+1:], blockStmt.List[insertIdx:])
									blockStmt.List[insertIdx] = assignToVoid
								}
							},
							Reset: func() {
								compositeLit.Elts = originalCompositeLitElements

								if assignToVoid != nil {
									if insertIdx < len(blockStmt.List)-1 {
										copy(blockStmt.List[insertIdx:], blockStmt.List[insertIdx+1:])
									}
									blockStmt.List[len(blockStmt.List)-1] = nil
									blockStmt.List = blockStmt.List[:len(blockStmt.List)-1]
								}
							},
						})
					}
				}
			}
		}
	}

	return mutations
}
