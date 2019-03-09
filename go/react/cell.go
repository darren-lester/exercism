package react

type cell struct {
	reactor    *reactor
	value      int
	dependants []*cell
	valueFn    func() int
	callbacks  map[*func(int)]bool
}

// Value returns the value of a cell
func (c *cell) Value() int {
	return c.value
}

// SetValue sets the value of a cell and triggers the reactor to update
func (c *cell) SetValue(v int) {
	c.setValue(v, true)
}

func (c *cell) setValue(v int, runCallbacks bool) {
	if v == c.value {
		return
	}
	prevValue := c.value
	c.value = v
	c.reactor.update(c, prevValue, runCallbacks)
}

// AddCallback adds a function to be called when a cell's value updates
func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks[&callback] = true
	return &canceler{cancel: func() {
		delete(c.callbacks, &callback)
	}}
}
