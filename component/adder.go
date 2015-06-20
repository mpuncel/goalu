package component

func HALFADDER(in1, in2 Outputer) (sum Outputer, cout Outputer){
  sum = XOR(in1.Output(), in2.Output())
  cout = AND(in1.Output(), in2.Output())

  return sum, cout
}

func FULLADDER(in1, in2, cin Outputer) (sum Outputer, cout Outputer){
  sum1, cout1 := HALFADDER(in1, in2)
  sum, cout2 := HALFADDER(sum1, cin)
  cout = XOR(cout2.Output(), cout1.Output())
  return sum, cout
}
