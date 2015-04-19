package component

import (
  "testing"
)

func TestOR(t *testing.T) {
  in1 := make(chan bool)
  in2 := make(chan bool)
  out := make(chan bool)

  or := OR{in1, in2, out}

  go func() {
    for {
      or.do()
    }
  }()
  in1 <- true
  in2 <- true
  if !<-out {
    t.Error("Expected true OR true to be true")
  }

  in1 <- true
  in2 <- false
  if !<-out {
    t.Error("Expected true OR false to be true")
  }

  in1 <- false
  in2 <- true
  if !<-out {
    t.Error("Expected false OR true to be true")
  }

  in1 <- false
  in2 <- false
  if <-out {
    t.Error("Expected false OR false to be false")
  }
}
