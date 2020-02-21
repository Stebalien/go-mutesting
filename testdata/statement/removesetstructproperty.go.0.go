// +build test

package example

func statementRemoveSetStructPropertyWhichIsAnotherStruct() Outer {
	outer := Outer{}
	_ = outer.inner

	return outer
}

type Inner struct {
}

type Outer struct {
	inner Inner
}
