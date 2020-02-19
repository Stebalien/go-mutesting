package expression

import (
	"go-mutesting/test"
	"testing"
)

func TestMutatorRemoveTerm(t *testing.T) {
	test.Mutator(
		t,
		MutatorRemoveTerm,
		"../../testdata/expression/remove.go",
		6,
	)
}
