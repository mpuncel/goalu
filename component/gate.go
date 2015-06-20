package component

func XOR(in1, in2 chan bool) Outputer{
  out := make(chan bool)

  go func() {
    for {
      val1 := <-in1
      val2 := <-in2

      out <- val1 != val2
    }
  }()

  return NODE(out)
}

func OR(in1, in2 chan bool) Outputer {
  out := make(chan bool)
  go func() {
    for {
      val1 := <-in1
      val2 := <-in2

      out <- val1 || val2
    }
  }()

  return NODE(out)
}

func AND(in1, in2 chan bool) Outputer{
  out := make(chan bool)
  go func() {
    for {
      val1 := <-in1
      val2 := <-in2

      out <- val1 && val2
    }
  }()

  return NODE(out)
}

