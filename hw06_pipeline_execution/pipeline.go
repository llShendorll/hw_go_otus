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
		stageOut := make(Bi)

		go func(out In, done In, stageOut Bi) {
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
		}(out, done, stageOut)

		out = stage(stageOut)
	}
	return out
}
