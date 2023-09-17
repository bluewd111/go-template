package functions

func MapValues[T comparable, U any](m map[T]*U) []*U {
	vs := []*U{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}
