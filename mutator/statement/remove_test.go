package statement

import (
	"testing"

	"github.com/Stebalien/go-mutesting/test"
)

func TestMutatorRemoveStatement(t *testing.T) {
	test.Mutator(
		t,
		MutatorRemoveStatement,
		"../../testdata/statement/remove.go",
		19,
	)
}
