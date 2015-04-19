package component

type xor struct {
  in1, in2, out chan bool
}

func XOR(in1, in2, out chan bool) {
  x := xor{in1, in2, out}

  go x.start()
}

func (x *xor) start() {
  for {
    val1 := <-x.in1
    val2 := <-x.in2

    x.out <- val1 != val2
  }
}

type or struct {
  in1, in2, out chan bool
}

func OR(in1, in2, out chan bool) {
  o := or{in1, in2, out}

  go o.start()
}

func (o *or) start() {
  for {
    val1 := <-o.in1
    val2 := <-o.in2

    o.out <- val1 || val2
  }
}

type and struct {
  in1, in2, out chan bool
}

func AND(in1, in2, out chan bool) {
  and := and{in1, in2, out}

  go and.start()
}

func (a *and) start() {
  for {
    val1 := <-a.in1
    val2 := <-a.in2

    a.out <- val1 && val2
  }
}

type adder struct {
  in1, in2, out, carry chan bool
}

func ADDER(in1, in2, out, carry chan bool) {
  go func() {
    for {
      val1 := <- in1
      val2 := <- in2

      xor1 := make(chan bool)
      xor2 := make(chan bool)

      XOR(xor1, xor2, out)
      xor1 <- val1
      xor2 <- val2

      and1 := make(chan bool)
      and2 := make(chan bool)
      AND(and1, and2, carry)

      and1 <- val1
      and2 <- val2
    }
  }()
}
