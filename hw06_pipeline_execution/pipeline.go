package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)
	finish := make(chan struct{})
	outs := prepareOuts(in, stages...)
	go RunOuts(outs, done, finish)

	go func() {
	handler:
		for {
			select {
			case res, ok := <-outs[len(outs)-1]:
				if !ok {
					close(finish)
					break handler
				}
				out <- res
			case <-done:
				break handler
			}
		}
		close(out)
	}()

	return out
}

func prepareOuts(in In, stages ...Stage) []Out {
	outs := make([]Out, len(stages))
	for i, stage := range stages {
		outs[i] = stage(in)
		in = outs[i]
	}
	return outs
}

func RunOuts(outs []Out, done In, finish chan struct{}) {
	select {
	case <-done:
		for _, currOut := range outs {
			go readUntilClosed(currOut)
		}
	case <-finish:
		return
	}
}

func readUntilClosed(out Out) {
	for range out {
	}
}
