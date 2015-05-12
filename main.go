package main

// install gen to get this to work
//     go get github.com/clipperhouse/gen
//
// run go generate to rebuild the *_slice.go files

//go:generate gen

// +gen slice:"Where,Count,GroupBy[string],GroupBy[int64]"
type MyType struct{}

// +gen slice:"GroupBy[MyType]"
type NewType struct{}

func main() {

}
