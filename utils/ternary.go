package utils

// Ternary implements normal ternary operator in golang (when in stdlib?)
// something = Ternary(someExpression, caseIfTrue, caseIfFalse)
func Ternary[K any](expression bool, caseIfTrue K, caseIfFalse K) K {
	if expression {
		return caseIfTrue
	} else {
		return caseIfFalse
	}
}
