package component

// broadcasts the value from an input channel to many listeners.
type node struct {
	input   chan bool
	outputs []chan bool
}

type Outputer interface {
	Output() chan bool
}

func NODE(in chan bool) Outputer {
	node := &node{
		input: in,
	}
	go node.run()
	return node
}

func (n *node) Output() chan bool {
	out := make(chan bool)
	n.outputs = append(n.outputs, out)
	return out
}

func (n *node) run() {
	for in := range n.input {
		for _, output := range n.outputs {
			output <- in
		}
	}
}
