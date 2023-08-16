package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	pipeLine := func(in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
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
					case out <- res:
					}
				}
			}
		}()
		return out
	}
	for _, stage := range stages {
		out := pipeLine(stage(in))
		in = out
	}
	return in
}
