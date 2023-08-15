package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	i := 0
	var pipelineFunc func(in In) Out
	pipelineFunc = func(in In) Out {
		out := make(Bi)
		i = i + 1
		go func(pos int) {
			defer close(out)
			for result := range stages[pos](in) {
				select {
				case <-done:
					return
				case out <- result:
				}
			}
		}(i - 1)
		if i == len(stages) {
			return out
		}
		return pipelineFunc(out)
	}
	return pipelineFunc(in)
}
