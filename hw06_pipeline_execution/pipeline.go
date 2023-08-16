package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	pipeLine := func(in In, stage Stage) Out {
		inp := make(Bi)
		outp := make(Bi)
		go func() {
			defer close(inp)
			for result := range in {
				select {
				case <-done:
					return
				case inp <- result:
				}
			}
		}()
		go func() {
			defer close(outp)
			for result := range stage(inp) {
				select {
				case <-done:
					return
				case outp <- result:
				}
			}
		}()
		return outp
	}

	for _, stage := range stages {
		out := pipeLine(in, stage)
		in = out
	}
	return in
}
