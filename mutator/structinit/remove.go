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
	var strus []*ast.CompositeLit
	_ = strus

	var mutations []mutator.Mutation

	switch n := node.(type) {
	case *ast.BlockStmt:
		for _, spec := range n.List {
			switch a := spec.(type) {
			case *ast.AssignStmt:
				if cl, ok := a.Rhs[0].(*ast.CompositeLit); ok {
					if _, ok := cl.Type.(*ast.Ident).Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType); ok {
						strus = append(strus, cl)
					}
				}

			}
		}
	default:
		return mutations
	}

	for i, stru := range strus {
		old := *stru

		//for j, elt := range strus[i].Elts {
		//
		//}

		mutations = append(mutations, mutator.Mutation{
			Change: func() {
				elts := []ast.Expr{stru.Elts[0]}
				strus[i].Elts = elts
			},
			Reset: func() {
				strus[i] = &old
			},
		})

	}

	//for i, _ := range strus {
	//	currentStruct, ok := strus[i].(*ast.TypeSpec).Type.(*ast.StructType)
	//	if !ok {
	//		continue
	//	}
	//
	//	oldStructFields := *strus[i].(*ast.TypeSpec).Type.(*ast.StructType).Fields
	//	_ = oldStructFields
	//
	//	for k, field := range oldStructFields.List {
	//		currentStruct.Fields.List = []*ast.Field{}
	//
	//		mutations = append(mutations, mutator.Mutation{
	//			Change: func() {
	//				for j := 0; j < len(oldStructFields.List); j++ {
	//					if k != j {
	//						currentStruct.Fields.List = append(currentStruct.Fields.List, field)
	//					}
	//				}
	//			},
	//			Reset: func() {
	//				strus[i].(*ast.TypeSpec).Type.(*ast.StructType).Fields = &oldStructFields
	//			},
	//		})
	//	}
	//}

	return mutations
}
