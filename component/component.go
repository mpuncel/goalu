package component

type OR struct {
  in1 chan bool
  in2 chan bool

  out chan bool
}

func (or *OR) do() {
  val1 := <-or.in1
  val2 := <-or.in2

  or.out <- val1 || val2
}
