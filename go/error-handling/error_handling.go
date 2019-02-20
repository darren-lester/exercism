package erratum

// Use opens a resource
func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return
	}

	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(FrobError); ok {
				resource.Defrob(v.defrobTag)
			}
			err = r.(error)
		}
		resource.Close()
	}()

	resource.Frob(input)

	return
}
