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
type NewType struct {
	m MyType
}

func main() {
	t := []MyType{
		{name: "woot"},
		{name: "moo"},
		{name: "rawr"},
	}

	log.Println("group by a string")
	grouped := MyTypeSlice(t).GroupByString(func(m MyType) string { return m.name })

	for name, elements := range grouped {
		log.Printf("name=%q elements=%v", name, elements)
	}

	log.Println("count specific elements")
	q := "rawr"
	count := MyTypeSlice(t).Count(func(m MyType) bool {
		if m.name == q {
			return true
		}
		return false
	})

	log.Printf("Found %d %q-s", count, q)

	nt := []NewType{
		{m: t[0]},
		{m: t[1]},
		{m: t[2]},
	}

	log.Print("grouping by a custom type")
	groupedByMyType := NewTypeSlice(nt).GroupByMyType(func(n NewType) MyType { return n.m })
	for myType, groups := range groupedByMyType {
		log.Printf("%+v: %+v", myType, groups)

	}
}
