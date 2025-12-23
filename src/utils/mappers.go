package utils

func MapSF[T comparable](o []T, extractor func(T) int64) []int64 {
	result := []int64{}
	for _, supplier := range o {
		result = append(result, extractor(supplier))
	}
	return result
}

func AssossiateBy[K comparable, V, T any](o []T, extractor func(T) (K, V)) map[K]V {
	retult := map[K]V{}
	for _, supplier := range o {
		key, value := extractor(supplier)
		retult[key] = value
	}
	return retult
}
func AssossiateSliceBy[K comparable, V, T any](o []T, extractor func(T) (K, V)) map[K][]V {
	retult := map[K][]V{}
	for _, supplier := range o {
		key, value := extractor(supplier)
		if _, ok := retult[key]; ok {
			retult[key] = append(retult[key], value)
			continue
		}
		retult[key] = []V{value}
	}
	return retult
}

func Map[R, T any](o []T, extractor func(T) R) []R {
	result := []R{}
	for _, supplier := range o {
		result = append(result, extractor(supplier))
	}
	return result
}
