package go_func

func Combine2[A, B, C any](func1 func(A) B, func2 func(B) C) func(A) C {
	return func(t A) C {
		return func2(func1(t))
	}
}
func Combine3[A, B, C, D any](func1 func(A) B, func2 func(B) C, func3 func(C) D) func(A) D {
	return func(t A) D {
		return func3(func2(func1(t)))
	}
}
func Combine4[A, B, C, D, E any](func1 func(A) B, func2 func(B) C, func3 func(C) D, func4 func(D) E) func(A) E {
	return func(t A) E {
		return func4(func3(func2(func1(t))))
	}
}
func Combine5[A, B, C, D, E, F any](func1 func(A) B, func2 func(B) C, func3 func(C) D, func4 func(D) E, func5 func(E) F) func(A) F {
	return func(t A) F {
		return func5(func4(func3(func2(func1(t)))))
	}
}
