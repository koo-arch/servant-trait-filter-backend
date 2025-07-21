package util

func ConvertSlice[T any, D any](items []T, convertFunc(func(T) D)) []D {
	dtos := make([]D, len(items))
	for i, item := range items {
		dtos[i] = convertFunc(item)
	}
	return dtos
}