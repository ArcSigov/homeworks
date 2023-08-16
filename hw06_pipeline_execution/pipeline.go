package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make([]Bi, len(stages))

	for i := range out {
		out[i] = make(Bi)
	}

	pipeLine := func(in In, pos int) {
		defer close(out[pos])
		for {
			select {
			case <-done:
				return
			case res, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case out[pos] <- res:
				}
			}
		}
	}
	if in != nil {
		for pos, stage := range stages {
			go pipeLine(stage(in), pos)
			in = out[pos]
		}
	} else {
		close(out[len(stages)-1])
	}
	return out[len(stages)-1]
}
