package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	resultStream := make(chan interface{})
	finished := make(chan struct{})
	go func() {
		defer close(resultStream)
		for {
			select {
			case <-done:
				return
			case <-finished:
				return
			}
		}
	}()
	go func() {
		var dataStream In
		for i, stage := range stages {
			switch i {
			case 0:
				dataStream = stageExec(done, in, stage)
			case len(stages) - 1:
				for val := range stageExec(done, dataStream, stage) {
					resultStream <- val
				}
				finished <- struct{}{}
			default:
				dataStream = stageExec(done, dataStream, stage)
			}
		}
	}()
	return resultStream
}

func stageExec(done In, in In, stage Stage) Out {
	stageStream := make(chan interface{})
	go func() {
		defer close(stageStream)
		for val := range stage(in) {
			select {
			case <-done:
				return
			default:
				stageStream <- val
			}
		}
	}()
	return stageStream
}
