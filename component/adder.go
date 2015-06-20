package component

func HALFADDER(in1, in2 Outputer) (sum Outputer, cout Outputer){
  sum = XOR(in1, in2)
  cout = AND(in1, in2)

  return sum, cout
}

func FULLADDER(in1, in2, cin Outputer) (sum Outputer, cout Outputer){
  sum1, cout1 := HALFADDER(in1, in2)
  sum, cout2 := HALFADDER(sum1, cin)
  cout = XOR(cout2, cout1)

  return sum, cout
}
