package structinit

import (
	"testing"

	"github.com/Stebalien/go-mutesting/test"
)

func TestMutatorInitStruct(t *testing.T) {
	test.Mutator(
		t,
		MutatorStructinitRemove,
		"../../testdata/structinit/remove.go",
		4,
	)
}
