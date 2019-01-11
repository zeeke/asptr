package asptr

// Int returns a pointer to the input value
func Int(v int) *int {
	return &v
}

// String returns a pointer to the input value
func String(v string) *string {
	return &v
}

// Bool returns a pointer to the input value
func Bool(v bool) *bool {
	return &v
}

// Float32 returns a pointer to the input value
func Float32(v float32) *float32 {
	return &v
}

// Float64 returns a pointer to the input value
func Float64(v float64) *float64 {
	return &v
}
