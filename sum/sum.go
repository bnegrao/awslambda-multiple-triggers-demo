package sum

type Operands struct {
	A int
	B int
}

func Sum(operands Operands) int {
	return operands.A + operands.B
}
