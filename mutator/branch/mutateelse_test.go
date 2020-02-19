package branch

import (
	"go-mutesting/test"
	"testing"
)

func TestMutatorElse(t *testing.T) {
	test.Mutator(
		t,
		MutatorElse,
		"../../testdata/branch/mutateelse.go",
		1,
	)
}
