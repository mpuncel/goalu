package component

import (
  "testing"
)

func TestXOR(t *testing.T) {
  in1 := make(chan bool)
  in2 := make(chan bool)
  out := make(chan bool)

  XOR(in1, in2, out)

  in1 <- true
  in2 <- true
  if <-out {
    t.Error("Expected 1 OR 1 to be 0")
  }

  in1 <- true
  in2 <- false
  if !<-out {
    t.Error("Expected 1 OR 0 to be 1")
  }

  in1 <- false
  in2 <- true
  if !<-out {
    t.Error("Expected 0 OR 1 to be 1")
  }

  in1 <- false
  in2 <- false
  if <-out {
    t.Error("Expected 0 OR 0 to be 0")
  }
}

func TestOR(t *testing.T) {
  in1 := make(chan bool)
  in2 := make(chan bool)
  out := make(chan bool)

  OR(in1, in2, out)

  in1 <- true
  in2 <- true
  if !<-out {
    t.Error("Expected 1 OR 1 to be 1")
  }

  in1 <- true
  in2 <- false
  if !<-out {
    t.Error("Expected 1 OR 0 to be 1")
  }

  in1 <- false
  in2 <- true
  if !<-out {
    t.Error("Expected 0 OR 1 to be 1")
  }

  in1 <- false
  in2 <- false
  if <-out {
    t.Error("Expected 0 OR 0 to be 0")
  }
}

func TestAND(t *testing.T) {
  in1 := make(chan bool)
  in2 := make(chan bool)
  out := make(chan bool)

  AND(in1, in2, out)

  in1 <- true
  in2 <- true
  if !<-out {
    t.Error("Expected 1 AND 1 to be 1")
  }

  in1 <- true
  in2 <- false
  if <-out {
    t.Error("Expected 1 AND 0 to be 0")
  }

  in1 <- false
  in2 <- true
  if <-out {
    t.Error("Expected 0 AND 1 to be 0")
  }

  in1 <- false
  in2 <- false
  if <-out {
    t.Error("Expected 0 AND 0 to be 0")
  }
}

func TestADDER(t *testing.T) {
  in1 := make(chan bool)
  in2 := make(chan bool)
  out := make(chan bool)
  carry := make(chan bool)

  ADDER(in1, in2, out, carry)

  in1 <- true
  in2 <- true
  if <- out || !<- carry {
    t.Error("Expected 1 + 1 to be 10")
  }

  in1 <- true
  in2 <- false
  if !<- out || <- carry {
    t.Error("Expected 1 + 0 to be 01")
  }

  in1 <- false
  in2 <- true
  if !<- out || <- carry {
    t.Error("Expected 0 + 1 to be 01")
  }

  in1 <- false
  in2 <- false
  if <- out || <- carry {
    t.Error("Expected 0 + 0 to be 00")
  }
}
