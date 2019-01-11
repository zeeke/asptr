package asptr

// Int returns a pointer to the input value
func Int(v int) *int {
	return &v
}

// Int8 returns a pointer to the input value
func Int8(v int8) *int8 {
	return &v
}

// Int16 returns a pointer to the input value
func Int16(v int16) *int16 {
	return &v
}

// Int32 returns a pointer to the input value
func Int32(v int32) *int32 {
	return &v
}

// Int64 returns a pointer to the input value
func Int64(v int64) *int64 {
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
