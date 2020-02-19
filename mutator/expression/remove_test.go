package expression

import (
	"testing"

	"github.com/AntonStoeckl/go-mutesting/test"
)

func TestMutatorRemoveTerm(t *testing.T) {
	test.Mutator(
		t,
		MutatorRemoveTerm,
		"../../testdata/expression/remove.go",
		6,
	)
}
