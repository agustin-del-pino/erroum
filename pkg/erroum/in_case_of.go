package erroum

func InCaseOfAny[T any](e error, r T, d T) T {
	if IsAny(e) {
		return r
	}

	return d
}

func InCaseOf[T any](e error, t error, r T, d T) T {
	if Is(e, t) {
		return r
	}

	return d
}

func InCaseOfSome[T any](e error, r T, d T, t ...error) T {
	if IsSome(e, t...) {
		return r
	}
	return d
}
