package statement

import (
	"testing"

	"github.com/AntonStoeckl/go-mutesting/test"
)

func TestMutatorRemoveStatement(t *testing.T) {
	test.Mutator(
		t,
		MutatorRemoveStatement,
		"../../testdata/statement/remove.go",
		17,
	)
}

func TestMutatorRemoveStatementSetStructProperty(t *testing.T) {
	test.Mutator(
		t,
		MutatorRemoveStatement,
		"../../testdata/statement/removesetstructproperty.go",
		1,
	)
}
