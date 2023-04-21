package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func (stage Stage) process(in In, done In) Bi {
	stageChannel := make(Bi)

	go func(stageChannel Bi, in Out) {
		defer close(stageChannel)

		for {
			select {
			case <-done:
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				stageChannel <- value
			}
		}
	}(stageChannel, in)

	return stageChannel
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		out := make(chan interface{})
		close(out)
		return out
	}

	for _, stage := range stages {
		in = stage(stage.process(in, done))
	}

	return in
}
