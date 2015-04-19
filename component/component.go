package component

func XOR(in1, in2, out chan bool) {
  go func() {
    for {
      val1 := <-in1
      val2 := <-in2

      out <- val1 != val2
    }
  }()
}

func OR(in1, in2, out chan bool) {
  go func() {
    for {
      val1 := <-in1
      val2 := <-in2

      out <- val1 || val2
    }
  }()
}

func AND(in1, in2, out chan bool) {
  go func() {
    for {
      val1 := <-in1
      val2 := <-in2

      out <- val1 && val2
    }
  }()
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
