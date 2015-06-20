package component

func XOR(in1, in2 Outputer) Outputer {
	out := make(chan bool)

	in1Ch := in1.Output()
	in2Ch := in2.Output()
	go func() {
		for {
			val1 := <-in1Ch
			val2 := <-in2Ch

			out <- val1 != val2
		}
	}()

	return NODE(out)
}

func OR(in1, in2 Outputer) Outputer {
	out := make(chan bool)
	in1Ch := in1.Output()
	in2Ch := in2.Output()
	go func() {
		for {
			val1 := <-in1Ch
			val2 := <-in2Ch

			out <- val1 || val2
		}
	}()

	return NODE(out)
}

func AND(in1, in2 Outputer) Outputer {
	out := make(chan bool)
	in1Ch := in1.Output()
	in2Ch := in2.Output()
	go func() {
		for {
			val1 := <-in1Ch
			val2 := <-in2Ch

			out <- val1 && val2
		}
	}()

	return NODE(out)
}
