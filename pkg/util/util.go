package util

// ToPtr converts the input to a pointer to itself
func ToPtr[T any](v T) *T {
	return &v
}
