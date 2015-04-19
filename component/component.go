package component

func XOR(in1, in2 chan bool) (chan bool){
  out := make(chan bool)

  go func() {
    for {
      val1 := <-in1
      val2 := <-in2

      out <- val1 != val2
    }
  }()

  return out
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

func AND(in1, in2 chan bool) (chan bool) {
  out := make(chan bool)
  go func() {
    for {
      val1 := <-in1
      val2 := <-in2

      out <- val1 && val2
    }
  }()

  return out
}

func ADDER(in1, in2 chan bool)  (chan bool, chan bool){
  xor1 := make(chan bool)
  xor2 := make(chan bool)

  out := XOR(xor1, xor2)
  and1 := make(chan bool)
  and2 := make(chan bool)
  carry := AND(and1, and2)

  go func() {
    for {
      val1 := <- in1
      val2 := <- in2

      xor1 <- val1
      xor2 <- val2

      and1 <- val1
      and2 <- val2
    }
  }()

  return out, carry
}
