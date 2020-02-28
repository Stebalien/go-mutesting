package structinit

import (
	"go/ast"
	"go/types"

	"github.com/AntonStoeckl/go-mutesting/mutator"
)

func init() {
	mutator.Register("structinit/remove", MutatorStructinitRemove)
}

// MutatorStructinitRemove implements a mutator to remove fields from struct initialisation.
func MutatorStructinitRemove(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	var compositeLits []*ast.CompositeLit
	var mutations []mutator.Mutation

	if blockStmt, ok := node.(*ast.BlockStmt); ok {
		for _, stmt := range blockStmt.List {
			if assignStmt, ok := stmt.(*ast.AssignStmt); ok {
				if compositeLit, ok := assignStmt.Rhs[0].(*ast.CompositeLit); ok {
					if ident, ok := compositeLit.Type.(*ast.Ident); ok {
						if typeSpec, ok := ident.Obj.Decl.(*ast.TypeSpec); ok {
							if _, ok := typeSpec.Type.(*ast.StructType); ok {
								compositeLits = append(compositeLits, compositeLit)
							}
						}
					}
				}
			}
		}
	}

	for _, compositeLit := range compositeLits {
		for compositeLitElementIdx := range compositeLit.Elts {
			originalCompositeLitElements := compositeLit.Elts
			var newCompositeLitElements []ast.Expr

			for i := 0; i < len(compositeLit.Elts); i++ {
				if compositeLitElementIdx != i {
					newCompositeLitElements = append(newCompositeLitElements, compositeLit.Elts[i])
				}
			}

			mutations = append(mutations, mutator.Mutation{
				Change: func() {
					compositeLit.Elts = newCompositeLitElements
				},
				Reset: func() {
					compositeLit.Elts = originalCompositeLitElements
				},
			})
		}
	}

	return mutations
}
