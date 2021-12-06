package memorandum

type Memoran struct {
	state int
}

func NewMemoran(value int) *Memoran {
	return &Memoran{value}
}

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{value: value}
}

func (n *Number) Double() {
	n.value = 2 * n.value
}

func (n *Number) Half() {
	n.value /= 2
}

func (n *Number) Value() int {
	return n.value
}

func (n *Number) CreateMemoran() *Memoran {
	return NewMemoran(n.value)
}

func (n *Number) ReinstateMemoran(memo *Memoran) {
	n.value = memo.state
}
