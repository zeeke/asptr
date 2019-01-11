package asptr

func ExampleInt() {
	var _ *int = Int(42)
}

func ExampleString() {
	var _ *string = String("Hello World!")
}

func ExampleBool() {
	var _ *bool = Bool(true)
}

func ExampleFloat32() {
	var _ *float32 = Float32(4.2)
}

func ExampleFloat64() {
	var _ *float64 = Float64(4.2)
}
