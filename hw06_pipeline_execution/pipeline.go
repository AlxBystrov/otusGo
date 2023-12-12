package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	resultStream := make(Bi)
	gen := make(Bi)

	// gen -> to be able to stop immediately
	go func() {
		defer close(gen)
		for val := range in {
			select {
			case <-done:
				return
			case gen <- val:
			}
		}
	}()

	dataStream := make(In)
	for i, stage := range stages {
		switch i {
		case 0:
			// use gen for first stage
			dataStream = stage(gen)
		default:
			dataStream = stage(dataStream)
		}
	}

	go func() {
		defer close(resultStream)
		for {
			select {
			case <-done:
				return
			case val, ok := <-dataStream:
				if !ok {
					return
				}
				resultStream <- val
			}
		}
	}()
	return resultStream
}
