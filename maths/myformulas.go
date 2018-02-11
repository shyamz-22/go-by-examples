package maths

func AlgebraAPlusBFormula(a, b int) int {
	return Multiply(a , a) + Multiply(b , b) + Multiply(2 ,  Multiply(a , b))
}

func AlgebraAMinusBFormula(a, b int) int {
	return Multiply(a , a) - Multiply(b , b) - Multiply(2 ,  Multiply(a , b))
}

func Multiply(a, b int) int {
	return a * b
}

