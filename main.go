package main

import "log"

// install gen to get this to work
//     go get github.com/clipperhouse/gen
//
// run go generate to rebuild the *_slice.go files

//go:generate gen

// +gen slice:"Where,Count,GroupBy[string],GroupBy[int64]"
type MyType struct {
	name string
}

// +gen slice:"GroupBy[MyType]"
type NewType struct{}

func main() {
	t := []MyType{
		{name: "woot"},
		{name: "moo"},
		{name: "rawr"},
	}

	// use generated group by command
	grouped := MyTypeSlice(t).GroupByString(func(m MyType) string { return m.name })

	for name, elements := range grouped {
		log.Printf("name=%q elements=%v", name, elements)
	}

	// count specific elements

	q := "rawr"
	count := MyTypeSlice(t).Count(func(m MyType) bool {
		if m.name == q {
			return true
		}
		return false
	})

	log.Printf("Found %d %q-s", count, q)
}
