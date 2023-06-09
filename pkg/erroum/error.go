package erroum

type err string

func (e *err) Error() string {
	return string(*e)
}

func New(d string) error {
	if d == "" {
		panic("empty errors are not admitted")
	}
	e := err(d)
	return &e
}

func From(c string, d string, e ...error) error {
	ds := New(d).Error()

	for _, er := range e {
		ds += c + er.Error()
	}
	
	return Merge(c, New(d), e...)
}

func Merge(c string, e error, ers ...error) error {
	if e == nil {
		panic("cannot merger from no error")
	}

	ds := e.Error()

	for _, er := range ers {
		if er == nil {
			continue
		}
		ds += c + er.Error()
	}

	return New(ds)
}
