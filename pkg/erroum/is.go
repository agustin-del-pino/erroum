package erroum

func Is(e error, t error) bool {
	return (e != nil && t != nil) && e == t
}

func IsAny(e error) bool {
	return e != nil
}

func IsSome(e error, t ...error) bool {
	if e == nil {
		return false
	}

	for _, er := range t {
		if er == nil {
			continue
		}
		if e == er {
			return true
		}
	}

	return false
}
