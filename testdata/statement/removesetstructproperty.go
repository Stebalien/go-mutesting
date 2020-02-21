// +build test

package example

func statementRemoveSetStructPropertyWhichIsAnotherStruct() Outer {
	outer := Outer{}
	outer.inner = Inner{}

	return outer
}

type Inner struct {
}

type Outer struct {
	inner Inner
}
