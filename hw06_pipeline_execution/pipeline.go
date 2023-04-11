package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		out = stage(getStage(out, done))
	}

	return out
}

func getStage(out In, done In) Out {
	stageOut := make(Bi)
	go func() {
		defer close(stageOut)
		for {
			select {
			case <-done:
				return
			case val, ok := <-out:
				if !ok {
					return
				}
				stageOut <- val
			}
		}
	}()

	return stageOut
}
