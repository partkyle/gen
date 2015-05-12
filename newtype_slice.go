// Generated by: gen
// TypeWriter: slice
// Directive: +gen on NewType

package main

// NewTypeSlice is a slice of type NewType. Use it where you would use []NewType.
type NewTypeSlice []NewType

// GroupByMyType groups elements into a map keyed by MyType. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv NewTypeSlice) GroupByMyType(fn func(NewType) MyType) map[MyType]NewTypeSlice {
	result := make(map[MyType]NewTypeSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
