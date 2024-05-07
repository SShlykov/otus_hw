package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	currentIn := in

	for _, stage := range stages {
		currentIn = stage(currentIn)
	}

	out := make(Bi)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case res, ok := <-currentIn:
				if !ok {
					return
				}
				select {
				case out <- res:
				case <-done:
					return
				}
			}
		}
	}()

	return out
}
