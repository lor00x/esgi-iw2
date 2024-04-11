package math

func Add(a, b int) int {
	if a > 0 && b > 0 {
		return a + b
	}
	return a + b
}

type DivideError struct {
	Message string
}

func (d DivideError) Error() string {
	return d.Message
}

var DivideByZero = DivideError{
	Message: "divide by zero",
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivideByZero
	}
	return a / b, nil
}
