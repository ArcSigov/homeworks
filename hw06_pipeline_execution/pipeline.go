package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Out)
	for _, stage := range stages {
		out = func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for result := range in {
					select {
					case <-done:
						return
					case out <- result:
					}
				}
			}()
			return out
		}(stage(in))
		in = out
	}
	return out
}
