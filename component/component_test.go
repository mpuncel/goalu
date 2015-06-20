package component

import (
  "testing"
)

func TestXOR(t *testing.T) {
  in1 := make(chan bool)
  in2 := make(chan bool)

  out := XOR(in1, in2).Output()

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

  out := OR(in1, in2).Output()

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

  out := AND(in1, in2).Output()

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

func TestHALFADDER(t *testing.T) {
  in1 := make(chan bool)
  in2 := make(chan bool)
  node1 := NODE(in1)
  node2 := NODE(in2)

  sumOutputer, coutOutputer := HALFADDER(node1, node2)
  sum := sumOutputer.Output()
  cout := coutOutputer.Output()

  in1 <- true
  in2 <- true
  if !<-cout || <-sum {
    t.Error("Expected 1 + 1 to be 10")
  }

  in1 <- true
  in2 <- false
  if <-cout || !<-sum {
    t.Error("Expected 1 + 0 to be 01")
  }

  in1 <- false
  in2 <- true
  if <-cout || !<-sum {
    t.Error("Expected 0 + 1 to be 01")
  }

  in1 <- false
  in2 <- false
  if <-cout || <-sum {
    t.Error("Expected 0 + 0 to be 00")
  }
}

func TestFULLADDER(t *testing.T) {
  in1 := make(chan bool)
  in2 := make(chan bool)
  cin := make(chan bool)
  nodeIn1 := NODE(in1)
  nodeIn2 := NODE(in2)
  nodeCin := NODE(cin)

  sumOutputer, coutOutputer := FULLADDER(nodeIn1, nodeIn2, nodeCin)
  sum := sumOutputer.Output()
  cout := coutOutputer.Output()

  in1 <- false
  in2 <- false
  cin <- false
  if <-cout || <-sum {
    t.Error("Expected 0 + 0 + 0 to be 00")
  }

  in1 <- false
  in2 <- false
  cin <- true
  if <-cout || !<-sum {
    t.Error("Expected 0 + 0 + 1 to be 01")
  }

  in1 <- false
  in2 <- true
  cin <- false
  if <-cout || !<-sum {
    t.Error("Expected 0 + 1 + 0 to be 01")
  }

  in1 <- false
  in2 <- true
  cin <- true
  if !<-cout || <-sum {
    t.Error("Expected 0 + 1 + 1 to be 10")
  }

  in1 <- true
  in2 <- false
  cin <- false
  if <-cout || !<-sum {
    t.Error("Expected 1 + 0 + 0 to be 01")
  }

  in1 <- true
  in2 <- false
  cin <- true
  if !<-cout || <-sum {
    t.Error("Expected 1 + 0 + 1 to be 10")
  }

  in1 <- true
  in2 <- true
  cin <- false
  if !<-cout || <-sum {
    t.Error("Expected 1 + 1 + 0 to be 10")
  }

  in1 <- true
  in2 <- true
  cin <- true
  if !<-cout || !<-sum {
    t.Error("Expected 1 + 1 + 1 to be 11")
  }
}
