package component

type orComponent struct {
  in1, in2, out chan bool
}

func OR(in1, in2, out chan bool) {
  or := orComponent{in1, in2, out}

  go func() {
    for {
      or.do()
    }
  }()
}

func (or *orComponent) do() {
  val1 := <-or.in1
  val2 := <-or.in2

  or.out <- val1 || val2
}
