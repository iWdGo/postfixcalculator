package RPN

import (
	"fmt"
	"testing"
)

/*
Running a single test
*/
func ExampleRPN_emulating_stack() {
	fmt.Printf("%.13f", RPN_emulating_stack(input))
	// Output: 3.0001220703125
}

func ExampleRPN_emulating_stack_sqrt() {
	fmt.Printf("%f", RPN_emulating_stack(RPNInput))
	// Output: 30.000000
}

func TestRPN_emulating_stack(t *testing.T) {
	if got := RPN_emulating_stack(input); got != input_want {
		t.Errorf("RPN(%s): got %f, want %f", input, got, input_want)
	}
}

//To check panicing on invalid expression
func TestRPN_emulating_stack_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	// The following is the code under test
	RPN_emulating_stack(InvalidInput)
}

func BenchmarkRPN_emulating_stack(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		RPN_emulating_stack(input)
	}
}
