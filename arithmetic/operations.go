package arithmetic

import "github.com/go-by-examples/maths"

//c = (a+b)^ 2 - a)
func Compute(a, b int) int{

	x := maths.AlgebraAPlusBFormula(a, b)
	y := x - a

	return y
}
