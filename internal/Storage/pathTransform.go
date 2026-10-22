package storage

type pathTransformer func(string) string


func DefaultPathTransform(key string) string {
	return key
}