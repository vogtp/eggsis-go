package thing

func (t *Thing) IsDead() bool {
	d := t.LP <= 0
	return d
}
