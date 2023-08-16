package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make([]Bi, len(stages))
	pipeLine := func(in In, pos int) {
		defer close(out[pos])
		for result := range in {
			select {
			case <-done:
				return
			case out[pos] <- result:
			}
		}
	}

	for pos, stage := range stages {
		out[pos] = make(Bi)
		go pipeLine(stage(in), pos)
		in = out[pos]
	}
	return out[len(stages)-1]
}
