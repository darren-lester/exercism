package react

type canceler struct {
	cancel func()
}

// Cancel cancels an action
func (c *canceler) Cancel() {
	c.cancel()
}
