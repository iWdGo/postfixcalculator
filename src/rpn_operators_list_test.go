package RPN

import (
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPN_operators_list() {
	fmt.Printf("%f", RPN_operators_list(RPNInput))
	// Output: 30.000000
}

func ExampleRPN_operators_listExp() {
	fmt.Printf("%.13f", RPN_operators_list(input))
	// Output: 3.0001220703125
}

func TestRPN_operators_list(t *testing.T) {
	if got := RPN_operators_list(RPNInput); got != RPNInput_want {
		t.Errorf("RPN(%s): got %f, want %f", RPNInput, got, RPNInput_want)
	}
}

func BenchmarkRPN_operators_list(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN_operators_list(input)
	}
}
