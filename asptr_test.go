package asptr

import "fmt"

// Basic usage of asptr
func ExampleUsage() {
	fmt.Println(*Int(42))
	fmt.Println(*String("Hello World!"))
	fmt.Println(*Bool(true))
	fmt.Println(*Float32(4.2))
	fmt.Println(*Float64(4.2))

	// Output:
	// 42
	// Hello World!
	// true
	// 4.2
	// 4.2
}
