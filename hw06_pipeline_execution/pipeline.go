package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// task := make(Bi)
	// out := stages[0](task)
	// go func() {
	// 	defer close(task)
	// 	for res := range in {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case task <- res:
	// 		}
	// 	}
	// }()

	// for i := 1; i < len(stages); i++ {
	// 	out = stages[i](out)
	// }
	// return out

	out := make(Out)
	for i := 0; i < len(stages); i++ {
		i := i
		out = func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for result := range stages[i](in) {
					select {
					case <-done:
						return
					case out <- result:
					}
				}
			}()
			return out
		}(in)
		in = out
	}
	return out
}
